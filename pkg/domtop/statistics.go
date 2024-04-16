package domtop

import (
	"fmt"
	"log/slog"

	"github.com/stanislav-zeman/domtop/pkg/statistics"
	"libvirt.org/go/libvirt"
)

func (dt *Domtop) cpuStats() (stats statistics.CPUs, err error) {
	domainCPUStats, err := dt.domain.GetCPUStats(-1, 1, 0)
	if err != nil {
		return
	}

	cpuStats := make([]statistics.CPU, 0, len(domainCPUStats))
	for _, cpuStat := range domainCPUStats {
		cpuStat := statistics.CPU{
			Time: cpuStat.CpuTime,
		}
		cpuStats = append(cpuStats, cpuStat)
	}

	stats = statistics.CPUs{
		CPUs: cpuStats,
	}
	return
}

func (dt *Domtop) diskStats() (stats statistics.Disks, err error) {
	blockStats, err := dt.domain.BlockStats("")
	if err != nil {
		err = fmt.Errorf("could not retrieve domain disk stats: %v", err)
		return
	}

	disk := statistics.Disk{
		Written: blockStats.WrBytes,
		Read:    blockStats.RdBytes,
	}
	stats = statistics.Disks{
		Disks: []statistics.Disk{
			disk,
		},
	}
	return
}

func (dt *Domtop) interfaceStats() (stats statistics.Interfaces, err error) {
	addresses, err := dt.domain.ListAllInterfaceAddresses(libvirt.DOMAIN_INTERFACE_ADDRESSES_SRC_AGENT)
	if err != nil {
		err = fmt.Errorf("could not retrieve domain interface addresses: %v", err)
		return
	}

	interfaceStats := make([]statistics.Interface, 0, len(addresses))
	for _, address := range addresses {
		domainInterfaceStats, err := dt.domain.InterfaceStats(address.Name)
		if err != nil {
			slog.Error("could not retrieve domain interface stats", "error", err, "address", address.Name)
			continue
		}

		interfaceStat := statistics.Interface{
			Rx:        domainInterfaceStats.RxBytes,
			Tx:        domainInterfaceStats.TxBytes,
			RxPackets: domainInterfaceStats.RxPackets,
			TxPackets: domainInterfaceStats.TxPackets,
		}
		interfaceStats = append(interfaceStats, interfaceStat)
	}

	stats = statistics.Interfaces{
		Interfaces: interfaceStats,
	}
	return
}
