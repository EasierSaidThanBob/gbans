package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/leighmacdonald/gbans/internal/errs"
	"github.com/leighmacdonald/gbans/pkg/util"
)

// Blocker provides a simple interface for blocking users connecting from banned IPs. It's designed to
// download list of, for example, VPN CIDR blocks, parse them and block any ip that is contained within any of those
// network blocks.
//
// IPs can be individually whitelisted if a remote/3rd party source cannot be changed.
type Blocker struct {
	cidrRx      *regexp.Regexp
	blocks      map[string][]*net.IPNet
	whitelisted map[int]*net.IPNet
	sync.RWMutex
}

func NewBlocker() *Blocker {
	return &Blocker{
		blocks:      make(map[string][]*net.IPNet),
		whitelisted: make(map[int]*net.IPNet),
		cidrRx:      regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(/(3[0-2]|2[0-9]|1[0-9]|[0-9]))?$`),
	}
}

func (b *Blocker) IsMatch(addr net.IP) (string, bool) {
	b.RLock()
	defer b.RUnlock()

	for _, whitelisted := range b.whitelisted {
		if whitelisted.Contains(addr) {
			return "", false
		}
	}

	for name, networks := range b.blocks {
		for _, block := range networks {
			if block.Contains(addr) {
				return name, true
			}
		}
	}

	return "", false
}

func (b *Blocker) RemoveSource(name string) {
	b.Lock()
	defer b.Unlock()

	delete(b.blocks, name)
}

func (b *Blocker) RemoveWhitelist(id int) {
	b.Lock()
	defer b.Unlock()

	delete(b.whitelisted, id)
}

func (b *Blocker) AddWhitelist(id int, network *net.IPNet) {
	b.Lock()
	defer b.Unlock()

	b.whitelisted[id] = network
}

func (b *Blocker) AddRemoteSource(ctx context.Context, name string, url string) (int64, error) {
	req, errReq := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if errReq != nil {
		return 0, errors.Join(errReq, errs.ErrCreateRequest)
	}

	client := util.NewHTTPClient()

	resp, errResp := client.Do(req)
	if errResp != nil {
		return 0, errors.Join(errResp, errs.ErrRequestPerform)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("%w: %d", errs.ErrRequestInvalidCode, resp.StatusCode)
	}

	bodyBytes, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return 0, errors.Join(errRead, errs.ErrResponseBody)
	}

	var blocks []*net.IPNet //nolint:prealloc

	for _, line := range strings.Split(string(bodyBytes), "\n") {
		trimmed := strings.TrimSpace(line)
		if !b.cidrRx.MatchString(trimmed) {
			continue
		}

		_, cidrBlock, errBlock := net.ParseCIDR(trimmed)
		if errBlock != nil {
			continue
		}

		blocks = append(blocks, cidrBlock)
	}

	b.Lock()
	b.blocks[name] = blocks
	b.Unlock()

	return int64(len(blocks)), nil
}
