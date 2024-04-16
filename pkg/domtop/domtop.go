package domtop

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/stanislav-zeman/domtop/pkg/statistics"
	"libvirt.org/go/libvirt"
)

type Domtop struct {
	libvirt      *libvirt.Connect
	domain       *libvirt.Domain
	refreshTimer *time.Ticker
	exporterChan chan<- statistics.Serializable
}

func New(config Config, exporterChan chan<- statistics.Serializable) (domtop *Domtop, err error) {
	connection, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		err = fmt.Errorf("could not connect to hypervisor: %v", err)
		return
	}

	domains, err := connection.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		return
	}

	domtop = &Domtop{
		libvirt:      connection,
		refreshTimer: time.NewTicker(config.RefreshPeriod),
		exporterChan: exporterChan,
	}
	for _, domain := range domains {
		name, err := domain.GetName()
		if err == nil {
			slog.Error("could not get domain name: %v", err)
			continue
		}

		if name == config.Domain {
			domtop.domain = &domain
			continue
		}

		domain.Free()
	}

	if domtop.domain == nil {
		err = fmt.Errorf("could not find specified domain '%s'", config.Domain)
		return
	}

	connection.DomainEventGraphicsRegister(domtop.domain, domtop.GraphicsCallback)
	connection.DomainEventLifecycleRegister(domtop.domain, domtop.LifecycleCallback)
	connection.DomainEventRebootRegister(domtop.domain, domtop.RebootCallback)
	return
}

func (dt *Domtop) Run(ctx context.Context) error {
	for {
		select {
		case <-dt.refreshTimer.C:
			dt.refresh()

		case <-ctx.Done():
			dt.Close()
			return nil
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
