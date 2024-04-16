package domtop

import (
	"time"

	"github.com/stanislav-zeman/domtop/pkg/statistics"
	"libvirt.org/go/libvirt"
)

func (dt *Domtop) LifecycleCallback(c *libvirt.Connect, d *libvirt.Domain, event *libvirt.DomainEventLifecycle) {
	e := statistics.Event{
		Type: statistics.LifecycleEvenType,
		Time: time.Now(),
		Parameters: map[string]any{
			"detail": event.Detail,
		},
	}
	dt.exporterChan <- e
}

func (dt *Domtop) GraphicsCallback(c *libvirt.Connect, d *libvirt.Domain, event *libvirt.DomainEventGraphics) {
	e := statistics.Event{
		Type: statistics.GraphicsEventType,
		Time: time.Now(),
		Parameters: map[string]any{
			"phase":  event.Phase,
			"local":  event.Local,
			"remote": event.Remote,
		},
	}
	dt.exporterChan <- e
}

func (dt *Domtop) RebootCallback(c *libvirt.Connect, d *libvirt.Domain) {
	e := statistics.Event{
		Type: statistics.RebootEventType,
		Time: time.Now(),
	}
	dt.exporterChan <- e
}
