package exporter

type Serializable interface {
	Serialize() (string, error)
}

var _ Serializable = ResourceUsage{}
var _ Serializable = LifecycleEvent{}
var _ Serializable = GraphicsEvent{}
var _ Serializable = RebootEvent{}

type ResourceUsage struct {
	CPUS    uint `json:"cpus,omitempty"`
	Discs   uint `json:"discs,omitempty"`
	Network uint `json:"network,omitempty"`
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
