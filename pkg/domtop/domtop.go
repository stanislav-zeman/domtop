package domtop

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/stanislav-zeman/domtop/pkg/config"
	"github.com/stanislav-zeman/domtop/pkg/statistics"
	"libvirt.org/go/libvirt"
)

type Domtop struct {
	libvirt      *libvirt.Connect
	domain       *libvirt.Domain
	refreshTimer *time.Ticker
	exporterChan chan<- statistics.Serializable
}

func New(cfg config.Config, exporterChan chan<- statistics.Serializable) (domtop *Domtop, err error) {
	slog.Info("connecting to libvirt hypervisor", "hypervisorURL", cfg.HypervisorURL)
	connection, err := libvirt.NewConnect(cfg.HypervisorURL)
	if err != nil {
		err = fmt.Errorf("could not connect to hypervisor: %v", err)
		return
	}

	domain, err := connection.LookupDomainByName(cfg.DomainName)
	if err != nil {
		err = fmt.Errorf("could not find specified domain '%s': %v", cfg.DomainName, err)
		return
	}

	connection.DomainEventGraphicsRegister(domain, domtop.GraphicsCallback)
	connection.DomainEventLifecycleRegister(domain, domtop.LifecycleCallback)
	connection.DomainEventRebootRegister(domain, domtop.RebootCallback)
	domtop = &Domtop{
		libvirt:      connection,
		domain:       domain,
		refreshTimer: time.NewTicker(cfg.RefreshPeriod),
		exporterChan: exporterChan,
	}
	return
}

func (dt *Domtop) Run(ctx context.Context) {
	slog.Info("started domtop")
	for {
		select {
		case <-dt.refreshTimer.C:
			dt.refresh()

		case <-ctx.Done():
			dt.Close()
			return
		}
	}
}

func (dt *Domtop) Close() {
	dt.domain.Free()
	dt.libvirt.Close()
}

func (dt *Domtop) refresh() {
	cpuStats, err := dt.cpuStats()
	if err != nil {
		slog.Error("could not retrieve domain CPU stats", "error", err)
		return
	}

	diskStats, err := dt.diskStats()
	if err != nil {
		slog.Error("could not retrieve domain disk stats", "error", err)
		return
	}

	interfaceStats, err := dt.interfaceStats()
	if err != nil {
		slog.Error("could not retrieve domain interface stats", "error", err)
		return
	}

	stats := statistics.Usage{
		CPUs:       cpuStats,
		Disks:      diskStats,
		Interfaces: interfaceStats,
	}
	dt.exporterChan <- stats
}
