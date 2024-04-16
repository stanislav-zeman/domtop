package domtop

import (
	"testing"
	"time"

	"github.com/stanislav-zeman/domtop/pkg/statistics"
	"libvirt.org/go/libvirt"
)

func TestDomtop_GraphicsCallback(t *testing.T) {
	type fields struct {
		libvirt      *libvirt.Connect
		domain       *libvirt.Domain
		refreshTimer *time.Ticker
		exporterChan chan<- statistics.Serializable
	}
	type args struct {
		c     *libvirt.Connect
		d     *libvirt.Domain
		event *libvirt.DomainEventGraphics
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dt := &Domtop{
				libvirt:      tt.fields.libvirt,
				domain:       tt.fields.domain,
				refreshTimer: tt.fields.refreshTimer,
				exporterChan: tt.fields.exporterChan,
			}
			dt.GraphicsCallback(tt.args.c, tt.args.d, tt.args.event)
		})
	}
}
