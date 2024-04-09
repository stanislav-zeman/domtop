package domtop

import (
	"context"
	"time"
)

type Domtop struct {
	refreshTimer *time.Ticker
}

func NewDomtop(domain string, refreshTime time.Duration) *Domtop {
	return &Domtop{
		refreshTimer: time.NewTicker(refreshTime),
	}
}

func (dt *Domtop) Run(ctx context.Context) error {
	for {
		select {
		case <-dt.refreshTimer.C:
			dt.refresh()
		case <-ctx.Done():
			return nil
		}
	}
}

func (dt *Domtop) refresh() {
	// TODO
}
