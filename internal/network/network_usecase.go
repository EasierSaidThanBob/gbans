package network

import (
	"context"
	"errors"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/leighmacdonald/gbans/internal/domain"
	"github.com/leighmacdonald/gbans/pkg/fp"
	"github.com/leighmacdonald/gbans/pkg/ip2location"
	"github.com/leighmacdonald/gbans/pkg/logparse"
	"github.com/leighmacdonald/steamid/v3/steamid"
	"go.uber.org/zap"
)

type networkUsecase struct {
	nr      domain.NetworkRepository
	bl      domain.BlocklistUsecase
	pu      domain.PersonUsecase
	blocker *Blocker
	log     *zap.Logger
	eb      *fp.Broadcaster[logparse.EventType, logparse.ServerEvent]
}

func NewNetworkUsecase(log *zap.Logger, broadcaster *fp.Broadcaster[logparse.EventType, logparse.ServerEvent],
	nr domain.NetworkRepository, bl domain.BlocklistUsecase, pu domain.PersonUsecase,
) domain.NetworkUsecase {
	return networkUsecase{
		log:     log.Named("network"),
		nr:      nr,
		bl:      bl,
		blocker: NewBlocker(),
		eb:      broadcaster,
		pu:      pu,
	}
}

func (u networkUsecase) Start(ctx context.Context) {
	log := u.log.Named("playerConnectionWriter")

	serverEventChan := make(chan logparse.ServerEvent)
	if errRegister := u.eb.Consume(serverEventChan, logparse.Connected); errRegister != nil {
		log.Warn("logWriter Tried to register duplicate reader channel", zap.Error(errRegister))

		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		case evt := <-serverEventChan:
			newServerEvent, ok := evt.Event.(logparse.ConnectedEvt)
			if !ok {
				continue
			}

			if newServerEvent.Address == "" {
				log.Warn("Empty Person message body, skipping")

				continue
			}

			parsedAddr := net.ParseIP(newServerEvent.Address)
			if parsedAddr == nil {
				log.Warn("Received invalid address", zap.String("addr", newServerEvent.Address))

				continue
			}

			// Maybe ignore these and wait for connect call to create?
			var person domain.Person
			if errPerson := u.pu.GetOrCreatePersonBySteamID(ctx, newServerEvent.SID, &person); errPerson != nil {
				log.Error("Failed to load Person", zap.Error(errPerson))

				continue
			}

			conn := domain.PersonConnection{
				IPAddr:      parsedAddr,
				SteamID:     newServerEvent.SID,
				PersonaName: strings.ToValidUTF8(newServerEvent.Name, "_"),
				CreatedOn:   newServerEvent.CreatedOn,
				ServerID:    evt.ServerID,
			}

			lCtx, cancel := context.WithTimeout(ctx, time.Second*5)
			if errChat := u.nr.AddConnectionHistory(lCtx, &conn); errChat != nil {
				log.Error("Failed to add connection history", zap.Error(errChat))
			}

			cancel()
		}
	}
}

func (u networkUsecase) LoadNetBlocks(ctx context.Context) error {
	sources, errSource := u.bl.GetCIDRBlockSources(ctx)
	if errSource != nil {
		return errors.Join(errSource, domain.ErrInitNetBlocks)
	}

	var total atomic.Int64

	waitGroup := sync.WaitGroup{}

	for _, source := range sources {
		if !source.Enabled {
			continue
		}

		waitGroup.Add(1)

		go func(src domain.CIDRBlockSource) {
			defer waitGroup.Done()

			count, errAdd := u.blocker.AddRemoteSource(ctx, src.Name, src.URL)
			if errAdd != nil {
				u.log.Error("Could not load remote source URL")
			}

			total.Add(count)
		}(source)
	}

	waitGroup.Wait()

	whitelists, errWhitelists := u.bl.GetCIDRBlockWhitelists(ctx)
	if errWhitelists != nil {
		if !errors.Is(errWhitelists, domain.ErrNoResult) {
			return errors.Join(errWhitelists, domain.ErrInitNetWhitelist)
		}
	}

	for _, whitelist := range whitelists {
		u.blocker.AddWhitelist(whitelist.CIDRBlockWhitelistID, whitelist.Address)
	}

	u.log.Info("Loaded cidr block lists",
		zap.Int64("cidr_blocks", total.Load()), zap.Int("whitelisted", len(whitelists)))

	return nil
}

func (u networkUsecase) AddConnectionHistory(ctx context.Context, conn *domain.PersonConnection) error {
	return u.nr.AddConnectionHistory(ctx, conn)
}

func (u networkUsecase) IsMatch(addr net.IP) (string, bool) {
	return u.blocker.IsMatch(addr)
}

func (u networkUsecase) AddWhitelist(id int, network *net.IPNet) {
	u.blocker.AddWhitelist(id, network)
}

func (u networkUsecase) RemoveWhitelist(id int) {
	u.blocker.RemoveWhitelist(id)
}

func (u networkUsecase) AddRemoteSource(ctx context.Context, name string, url string) (int64, error) {
	return u.blocker.AddRemoteSource(ctx, name, url)
}

func (u networkUsecase) GetASNRecordByIP(ctx context.Context, ipAddr net.IP, asnRecord *ip2location.ASNRecord) error {
	return u.nr.GetASNRecordByIP(ctx, ipAddr, asnRecord)
}

func (u networkUsecase) GetASNRecordsByNum(ctx context.Context, asNum int64) (ip2location.ASNRecords, error) {
	return u.nr.GetASNRecordsByNum(ctx, asNum)
}

func (u networkUsecase) GetLocationRecord(ctx context.Context, ipAddr net.IP, record *ip2location.LocationRecord) error {
	return u.nr.GetLocationRecord(ctx, ipAddr, record)
}

func (u networkUsecase) GetProxyRecord(ctx context.Context, ipAddr net.IP, proxyRecord *ip2location.ProxyRecord) error {
	return u.nr.GetProxyRecord(ctx, ipAddr, proxyRecord)
}

func (u networkUsecase) InsertBlockListData(ctx context.Context, log *zap.Logger, blockListData *ip2location.BlockListData) error {
	return u.nr.InsertBlockListData(ctx, log, blockListData)
}

func (u networkUsecase) GetPersonIPHistory(ctx context.Context, sid64 steamid.SID64, limit uint64) (domain.PersonConnections, error) {
	return u.nr.GetPersonIPHistory(ctx, sid64, limit)
}

func (u networkUsecase) GetPlayerMostRecentIP(ctx context.Context, steamID steamid.SID64) net.IP {
	return u.nr.GetPlayerMostRecentIP(ctx, steamID)
}

func (u networkUsecase) QueryConnectionHistory(ctx context.Context, opts domain.ConnectionHistoryQueryFilter) ([]domain.PersonConnection, int64, error) {
	return u.nr.QueryConnectionHistory(ctx, opts)
}
