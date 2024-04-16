package domtop

import (
	"fmt"
	"log/slog"

	"github.com/stanislav-zeman/domtop/pkg/statistics"
	"libvirt.org/go/libvirt"
)

func (dt *Domtop) cpuStats() (stats []statistics.CPU, err error) {
	domainCPUStats, err := dt.domain.GetCPUStats(-1, 1, 0)
	if err != nil {
		return
	}

	stats = make([]statistics.CPU, 0, len(domainCPUStats))
	for _, cpuStat := range domainCPUStats {
		cpuStat := statistics.CPU{
			Time: cpuStat.CpuTime,
		}
		stats = append(stats, cpuStat)
	}

	return
}

func (dt *Domtop) diskStats() (stats []statistics.Disk, err error) {
	blockStats, err := dt.domain.BlockStats("")
	if err != nil {
		err = fmt.Errorf("could not retrieve domain disk stats: %v", err)
		return
	}

	disk := statistics.Disk{
		WrittenBytes: blockStats.WrBytes,
		ReadBytes:    blockStats.RdBytes,
	}
	stats = []statistics.Disk{
		disk,
	}
	return
}

func (dt *Domtop) interfaceStats() (stats []statistics.Interface, err error) {
	addresses, err := dt.domain.ListAllInterfaceAddresses(libvirt.DOMAIN_INTERFACE_ADDRESSES_SRC_AGENT)
	if err != nil {
		err = fmt.Errorf("could not retrieve domain interface addresses: %v", err)
		return
	}

	stats = make([]statistics.Interface, 0, len(addresses))
	for _, address := range addresses {
		domainInterfaceStats, err := dt.domain.InterfaceStats(address.Name)
		if err != nil {
			slog.Error("could not retrieve domain interface stats", "error", err, "address", address.Name)
			continue
		}

		interfaceStat := statistics.Interface{
			RxBytes:   domainInterfaceStats.RxBytes,
			TxBytes:   domainInterfaceStats.TxBytes,
			RxPackets: domainInterfaceStats.RxPackets,
			TxPackets: domainInterfaceStats.TxPackets,
		}
		stats = append(stats, interfaceStat)
	}

	return
}
