//go:build !baremetal
// +build !baremetal

// Some documentation for the BlueZ D-Bus interface:
// https://git.kernel.org/pub/scm/bluetooth/bluez.git/tree/doc

package bluetooth

import (
	"errors"

	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
)

type Adapter struct {
	Adapter              *adapter.Adapter1
	Id                   string
	cancelChan           chan struct{}
	defaultAdvertisement *Advertisement

	connectHandler func(device Address, connected bool)
}

// DefaultAdapter is the default adapter on the system. On Linux, it is the
// first adapter available.
//
// Make sure to call Enable() before using it to initialize the adapter.
var DefaultAdapter = &Adapter{
	connectHandler: func(device Address, connected bool) {
		return
	},
}

// Enable configures the BLE stack. It must be called before any
// Bluetooth-related calls (unless otherwise indicated).
func (a *Adapter) Enable() (err error) {
	if a.Id == "" {
		a.Adapter, err = api.GetDefaultAdapter()
		if err != nil {
			return
		}
		a.Id, err = a.Adapter.GetAdapterID()
	}
	return nil
}

func (a *Adapter) Address() (MACAddress, error) {
	if a.Adapter == nil {
		return MACAddress{}, errors.New("adapter not enabled")
	}
	mac, err := ParseMAC(a.Adapter.Properties.Address)
	if err != nil {
		return MACAddress{}, err
	}
	return MACAddress{MAC: mac}, nil
}
