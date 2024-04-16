package statistics

type Serializable interface {
	Serialize() (string, error)
}

var _ Serializable = Usage{}

type Usage struct {
	CPUs       []CPU       `json:"cpus,omitempty"`
	Disks      []Disk      `json:"disks,omitempty"`
	Interfaces []Interface `json:"interfaces,omitempty"`
}

type CPU struct {
	Time uint64 `json:"time,omitempty"`
}

type Disk struct {
	WrittenBytes int64 `json:"writen_bytes,omitempty"`
	ReadBytes    int64 `json:"read_bytes,omitempty"`
}

type Interface struct {
	RxBytes   int64 `json:"rx_bytes,omitempty"`
	TxBytes   int64 `json:"tx_bytes,omitempty"`
	RxPackets int64 `json:"rx_packets,omitempty"`
	TxPackets int64 `json:"tx_packets,omitempty"`
}

func (u Usage) Serialize() (string, error) {
	panic("unimplemented")
}
