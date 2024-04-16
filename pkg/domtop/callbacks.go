package domtop

import "libvirt.org/go/libvirt"

func (dt *Domtop) LifecycleCallback(c *libvirt.Connect, d *libvirt.Domain, event *libvirt.DomainEventLifecycle) {

}

func (dt *Domtop) GraphicsCallback(c *libvirt.Connect, d *libvirt.Domain, event *libvirt.DomainEventGraphics) {

}

func (dt *Domtop) RebootCallback(c *libvirt.Connect, d *libvirt.Domain) {

}
