package statistics

type Serializable interface {
	Serialize() (string, error)
}

var _ Serializable = ResourceUsage{}
var _ Serializable = LifecycleEvent{}
var _ Serializable = GraphicsEvent{}
var _ Serializable = RebootEvent{}

type ResourceUsage struct {
	CPUs       CPUs `json:"CPU,omitempty"`
	Disks      Disks
	Interfaces Interfaces
}

type CPUs struct {
	CPUs []CPU `json:"CPU Times,omitempty"`
}

type CPU struct {
	Time uint64
}

type Disks struct {
	Disks []Disk
}

type Disk struct {
	Written int64
	Read    int64
}

type Interfaces struct {
	Interfaces []Interface
}

type Interface struct {
	Rx        int64
	Tx        int64
	RxPackets int64
	TxPackets int64
}

func (r ResourceUsage) Serialize() (string, error) {
	panic("unimplemented")
}

type LifecycleEvent struct {
}

func (l LifecycleEvent) Serialize() (string, error) {
	panic("unimplemented")
}

type GraphicsEvent struct {
}

func (g GraphicsEvent) Serialize() (string, error) {
	panic("unimplemented")
}

type RebootEvent struct {
}

func (r RebootEvent) Serialize() (string, error) {
	panic("unimplemented")
}
