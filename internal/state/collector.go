package state

import (
	"context"
	"errors"
	"net"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/leighmacdonald/gbans/internal/domain"
	"github.com/leighmacdonald/gbans/pkg/ip2location"
	"github.com/leighmacdonald/rcon/rcon"
	"github.com/leighmacdonald/steamid/v3/extra"
	"github.com/leighmacdonald/steamweb/v2"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

var (
	ErrRCONCommand      = errors.New("failed to execute rcon command")
	ErrFailedToDialRCON = errors.New("failed to connect to conf")
)

type Collector struct {
	log              *zap.Logger
	statusUpdateFreq time.Duration
	msListUpdateFreq time.Duration
	updateTimeout    time.Duration
	masterServerList []serverLocation
	serverState      map[int]domain.ServerState
	stateMu          *sync.RWMutex
	configs          []domain.ServerConfig
	configMu         *sync.RWMutex
	maxPlayersRx     *regexp.Regexp
	serverUsecase    domain.ServersUsecase
	configUsecase    domain.ConfigUsecase
}

func NewCollector(logger *zap.Logger, serverUsecase domain.ServersUsecase) *Collector {
	const (
		statusUpdateFreq = time.Second * 20
		msListUpdateFreq = time.Minute * 5
		updateTimeout    = time.Second * 5
	)

	return &Collector{
		log:              logger,
		statusUpdateFreq: statusUpdateFreq,
		msListUpdateFreq: msListUpdateFreq,
		updateTimeout:    updateTimeout,
		serverState:      map[int]domain.ServerState{},
		stateMu:          &sync.RWMutex{},
		configMu:         &sync.RWMutex{},
		maxPlayersRx:     regexp.MustCompile(`^"sv_visiblemaxplayers" = "(\d{1,2})"\s`),
		serverUsecase:    serverUsecase,
	}
}

func (c *Collector) Configs() []domain.ServerConfig {
	c.configMu.RLock()
	defer c.configMu.RUnlock()

	var conf []domain.ServerConfig

	conf = append(conf, c.configs...)

	return conf
}

func (c *Collector) ExecRaw(ctx context.Context, addr string, password string, cmd string) (string, error) {
	conn, errConn := rcon.Dial(ctx, addr, password, time.Second*5)
	if errConn != nil {
		return "", errors.Join(errConn, ErrFailedToDialRCON)
	}

	resp, errExec := conn.Exec(cmd)
	if errExec != nil {
		return "", errors.Join(errExec, ErrRCONExecCommand)
	}

	if errClose := conn.Close(); errClose != nil {
		c.log.Error("Could not close rcon connection", zap.Error(errClose))
	}

	return resp, nil
}

func (c *Collector) GetServer(serverID int) (domain.ServerConfig, error) {
	c.configMu.RLock()
	c.configMu.RUnlock()

	configs := c.Configs()

	serverIdx := slices.IndexFunc(configs, func(serverConfig domain.ServerConfig) bool {
		return serverConfig.ServerID == serverID
	})

	if serverIdx == -1 {
		return domain.ServerConfig{}, domain.ErrUnknownServerID
	}

	return configs[serverIdx], nil
}

func (c *Collector) Current() []domain.ServerState {
	c.stateMu.RLock()
	defer c.stateMu.RUnlock()

	var curState []domain.ServerState //nolint:prealloc
	for _, s := range c.serverState {
		curState = append(curState, s)
	}

	sort.SliceStable(curState, func(i, j int) bool {
		return curState[i].Name < curState[j].Name
	})

	return curState
}

func (c *Collector) startMSL(ctx context.Context) {
	var (
		log            = c.log.Named("msl_update")
		mlUpdateTicker = time.NewTicker(c.msListUpdateFreq)
	)

	for {
		select {
		case <-mlUpdateTicker.C:
			newMsl, errUpdateMsl := c.updateMSL(ctx)
			if errUpdateMsl != nil {
				log.Error("Failed to update master server list", zap.Error(errUpdateMsl))

				continue
			}

			c.stateMu.Lock()
			c.masterServerList = newMsl
			c.stateMu.Unlock()
		case <-ctx.Done():
			return
		}
	}
}

func (c *Collector) onStatusUpdate(conf domain.ServerConfig, newState extra.Status, maxVisible int) {
	c.stateMu.Lock()
	defer c.stateMu.Unlock()

	server := c.serverState[conf.ServerID]
	server.PlayerCount = newState.PlayersCount

	if maxVisible >= 0 {
		server.MaxPlayers = maxVisible
	} else {
		server.MaxPlayers = newState.PlayersMax
	}

	if newState.ServerName != "" {
		server.Name = newState.ServerName
	}

	server.Version = newState.Version
	server.Edicts = newState.Edicts
	server.Tags = newState.Tags

	if newState.Map != "" && newState.Map != server.Map {
		server.Map = newState.Map
	}

	server.Players = newState.Players

	c.serverState[conf.ServerID] = server
}

func (c *Collector) setServerConfigs(configs []domain.ServerConfig) {
	c.configMu.Lock()
	defer c.configMu.Unlock()

	var gone []domain.ServerConfig

	for _, exist := range c.configs {
		exists := false

		for _, newConf := range configs {
			if exist.ServerID == newConf.ServerID {
				exists = true

				break
			}
		}

		if !exists {
			gone = append(gone, exist)
		}
	}

	for _, conf := range gone {
		delete(c.serverState, conf.ServerID)
	}

	c.stateMu.Lock()

	for _, cfg := range configs {
		if _, found := c.serverState[cfg.ServerID]; !found {
			addr, errResolve := ResolveIP(cfg.Host)
			if errResolve != nil {
				c.log.Warn("Failed to resolve server ip", zap.String("addr", addr), zap.Error(errResolve))
				addr = cfg.Host
			}

			c.serverState[cfg.ServerID] = domain.ServerState{
				ServerID:      cfg.ServerID,
				Name:          cfg.DefaultHostname,
				NameShort:     cfg.Tag,
				Host:          cfg.Host,
				Port:          cfg.Port,
				RconPassword:  cfg.RconPassword,
				ReservedSlots: cfg.ReservedSlots,
				CC:            cfg.CC,
				Region:        cfg.Region,
				Latitude:      cfg.Latitude,
				Longitude:     cfg.Longitude,
				IP:            addr,
			}

		}
	}

	c.stateMu.Unlock()

	c.configMu.Lock()
	c.configs = configs
	c.configMu.Unlock()
}

func (c *Collector) Update(serverID int, update domain.PartialStateUpdate) error {
	c.stateMu.Lock()
	defer c.stateMu.Unlock()

	curState, ok := c.serverState[serverID]
	if !ok {
		return domain.ErrUnknownServer
	}

	if update.Hostname != "" {
		curState.Name = update.Hostname
	}

	curState.Map = update.CurrentMap
	curState.PlayerCount = update.PlayersReal
	curState.MaxPlayers = update.PlayersVisible
	curState.Bots = update.PlayersTotal - update.PlayersReal
	c.serverState[serverID] = curState

	return nil
}

var (
	ErrStatusParse       = errors.New("failed to parse status response")
	ErrMaxPlayerIntParse = errors.New("failed to cast max players value")
	ErrMaxPlayerParse    = errors.New("failed to parse sv_visiblemaxplayers response")
	ErrServerListRequest = errors.New("failed to fetch updated list")
	ErrDNSResolve        = errors.New("failed to resolve server dns")
	ErrRCONExecCommand   = errors.New("failed to perform command")
)

func (c *Collector) status(ctx context.Context, serverID int) (extra.Status, error) {
	server, errServerID := c.GetServer(serverID)
	if errServerID != nil {
		return extra.Status{}, errServerID
	}

	statusResp, errStatus := c.ExecRaw(ctx, server.Addr(), server.RconPassword, "status")
	if errStatus != nil {
		return extra.Status{}, errStatus
	}

	status, errParse := extra.ParseStatus(statusResp, true)
	if errParse != nil {
		return extra.Status{}, errors.Join(errParse, ErrStatusParse)
	}

	return status, nil
}

const maxPlayersSupported = 101

func (c *Collector) maxVisiblePlayers(ctx context.Context, serverID int) (int, error) {
	server, errServerID := c.GetServer(serverID)
	if errServerID != nil {
		return 0, errServerID
	}

	maxPlayersResp, errMaxPlayers := c.ExecRaw(ctx, server.Addr(), server.RconPassword, "sv_visiblemaxplayers")
	if errMaxPlayers != nil {
		return 0, errMaxPlayers
	}

	matches := c.maxPlayersRx.FindStringSubmatch(maxPlayersResp)
	if matches == nil || len(matches) != 2 {
		return 0, ErrMaxPlayerParse
	}

	maxPlayers, errCast := strconv.ParseInt(matches[1], 10, 32)
	if errCast != nil {
		return 0, errors.Join(errCast, ErrMaxPlayerIntParse)
	}

	if maxPlayers > maxPlayersSupported {
		maxPlayers = -1
	}

	return int(maxPlayers), nil
}

func (c *Collector) startStatus(ctx context.Context) {
	var (
		logger             = c.log.Named("statusUpdate")
		statusUpdateTicker = time.NewTicker(c.statusUpdateFreq)
	)

	for {
		select {
		case <-statusUpdateTicker.C:
			waitGroup := &sync.WaitGroup{}
			successful := atomic.Int32{}
			existing := atomic.Int32{}

			c.stateMu.RLock()
			configs := c.configs
			c.stateMu.RUnlock()

			startTIme := time.Now()

			for _, serverConfigInstance := range configs {
				waitGroup.Add(1)

				go func(conf domain.ServerConfig) {
					defer waitGroup.Done()

					log := logger.Named(conf.Tag)

					status, errStatus := c.status(ctx, conf.ServerID)
					if errStatus != nil {
						return
					}

					maxVisible, errMaxVisible := c.maxVisiblePlayers(ctx, conf.ServerID)
					if errMaxVisible != nil {
						log.Warn("Got invalid max players value", zap.Error(errMaxVisible))
					}

					c.onStatusUpdate(conf, status, maxVisible)

					successful.Add(1)
				}(serverConfigInstance)
			}

			waitGroup.Wait()

			logger.Debug("RCON update cycle complete",
				zap.Int32("success", successful.Load()),
				zap.Int32("existing", existing.Load()),
				zap.Int32("fail", int32(len(configs))-successful.Load()),
				zap.Duration("duration", time.Since(startTIme)))
		case <-ctx.Done():
			return
		}
	}
}

func (c *Collector) updateMSL(ctx context.Context) ([]serverLocation, error) {
	allServers, errServers := steamweb.GetServerList(ctx, map[string]string{
		"appid":     "440",
		"dedicated": "1",
	})

	if errServers != nil {
		return nil, errors.Join(errServers, ErrServerListRequest)
	}

	var ( //nolint:prealloc
		communityServers []serverLocation
		stats            = newGlobalTF2Stats()
	)

	for _, base := range allServers {
		server := serverLocation{
			LatLong: ip2location.LatLong{},
			Server:  base,
		}

		stats.ServersTotal++
		stats.Players += server.Players
		stats.Bots += server.Bots

		switch {
		case server.MaxPlayers > 0 && server.Players >= server.MaxPlayers:
			stats.CapacityFull++
		case server.Players == 0:
			stats.CapacityEmpty++
		default:
			stats.CapacityPartial++
		}

		if server.Secure {
			stats.Secure++
		}

		region := SteamRegionIDString(SvRegion(server.Region))

		_, regionFound := stats.Regions[region]
		if !regionFound {
			stats.Regions[region] = 0
		}

		stats.Regions[region] += server.Players

		mapType := guessMapType(server.Map)

		_, mapTypeFound := stats.MapTypes[mapType]
		if !mapTypeFound {
			stats.MapTypes[mapType] = 0
		}

		stats.MapTypes[mapType]++
		if strings.Contains(server.GameType, "valve") ||
			!server.Dedicated ||
			!server.Secure {
			stats.ServersCommunity++

			continue
		}

		communityServers = append(communityServers, server)
	}

	return communityServers, nil
}

// todo add external
// conf := c.configUsecase.Config()
// if conf.Debug.AddRCONLogAddress != "" {
// c.LogAddressAdd(ctx, conf.Debug.AddRCONLogAddress)
// }

func (c *Collector) Start(ctx context.Context) {
	var (
		log          = c.log.Named("State")
		trigger      = make(chan any)
		updateTicker = time.NewTicker(time.Minute * 30)
	)

	go c.startMSL(ctx)
	go c.startStatus(ctx)

	go func() {
		trigger <- true
	}()

	for {
		select {
		case <-updateTicker.C:
			trigger <- true
		case <-trigger:
			servers, _, errServers := c.serverUsecase.GetServers(ctx, domain.ServerQueryFilter{
				QueryFilter:     domain.QueryFilter{Deleted: false},
				IncludeDisabled: false,
			})
			if errServers != nil && !errors.Is(errServers, domain.ErrNoResult) {
				log.Error("Failed to fetch servers, cannot update State", zap.Error(errServers))

				continue
			}

			var configs []domain.ServerConfig
			for _, server := range servers {
				configs = append(configs, newServerConfig(
					server.ServerID,
					server.ShortName,
					server.Name,
					server.Address,
					server.Port,
					server.RCON,
					server.ReservedSlots,
					server.CC,
					server.Region,
					server.Latitude,
					server.Longitude,
				))
			}

			c.setServerConfigs(configs)
		case <-ctx.Done():
			return
		}
	}
}

func newServerConfig(serverID int, name string, defaultHostname string, address string,
	port int, rconPassword string, reserved int, countryCode string, region string, lat float64, long float64,
) domain.ServerConfig {
	return domain.ServerConfig{
		ServerID:        serverID,
		Tag:             name,
		DefaultHostname: defaultHostname,
		Host:            address,
		Port:            port,
		RconPassword:    rconPassword,
		ReservedSlots:   reserved,
		CC:              countryCode,
		Region:          region,
		Latitude:        lat,
		Longitude:       long,
	}
}

type SvRegion int

const (
	RegionNaEast SvRegion = iota
	RegionNAWest
	RegionSouthAmerica
	RegionEurope
	RegionAsia
	RegionAustralia
	RegionMiddleEast
	RegionAfrica
	RegionWorld SvRegion = 255
)

func SteamRegionIDString(region SvRegion) string {
	switch region {
	case RegionNaEast:
		return "ne"
	case RegionNAWest:
		return "nw"
	case RegionSouthAmerica:
		return "sa"
	case RegionEurope:
		return "eu"
	case RegionAsia:
		return "as"
	case RegionAustralia:
		return "au"
	case RegionMiddleEast:
		return "me"
	case RegionAfrica:
		return "af"
	case RegionWorld:
		fallthrough
	default:
		return "wd"
	}
}

func guessMapType(mapName string) string {
	mapName = strings.TrimPrefix(mapName, "workshop/")
	pieces := strings.SplitN(mapName, "_", 2)

	if len(pieces) == 1 {
		return "unknown"
	}

	return strings.ToLower(pieces[0])
}

type globalTF2StatsSnapshot struct {
	StatID           int64          `json:"stat_id"`
	Players          int            `json:"players"`
	Bots             int            `json:"bots"`
	Secure           int            `json:"secure"`
	ServersCommunity int            `json:"servers_community"`
	ServersTotal     int            `json:"servers_total"`
	CapacityFull     int            `json:"capacity_full"`
	CapacityEmpty    int            `json:"capacity_empty"`
	CapacityPartial  int            `json:"capacity_partial"`
	MapTypes         map[string]int `json:"map_types"`
	Regions          map[string]int `json:"regions"`
	CreatedOn        time.Time      `json:"created_on"`
}

// func (stats globalTF2StatsSnapshot) trimMapTypes() map[string]int {
//	const minSize = 5
//
//	out := map[string]int{}
//
//	for keyKey, value := range stats.MapTypes {
//		mapKey := keyKey
//		if value < minSize {
//			mapKey = "unknown"
//		}
//
//		out[mapKey] = value
//	}
//
//	return out
// }

func newGlobalTF2Stats() globalTF2StatsSnapshot {
	return globalTF2StatsSnapshot{
		MapTypes:  map[string]int{},
		Regions:   map[string]int{},
		CreatedOn: time.Now(),
	}
}

type serverLocation struct {
	ip2location.LatLong
	steamweb.Server
}

func ResolveIP(addr string) (string, error) {
	ipAddr := net.ParseIP(addr)
	if ipAddr != nil {
		return ipAddr.String(), nil
	}

	ips, err := net.LookupIP(addr)
	if err != nil || len(ips) == 0 {
		return "", errors.Join(err, ErrDNSResolve)
	}

	return ips[0].String(), nil
}