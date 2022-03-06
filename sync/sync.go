//
// This file is part of quarto-discovery.
//
// Copyright 2018-2021 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to modify or
// otherwise use the software for commercial activities involving the Arduino
// software without disclosing the source code of your own applications. To purchase
// a commercial license, send an email to license@arduino.cc.
//

package sync

import (
	"github.com/arduino/go-properties-orderedmap"
	discovery "github.com/ben-qnimble/pluggable-discovery-protocol-handler/v3"
	"github.com/ben-qnimble/go-serial/enumerator"
)

// processUpdates sends 'add' and 'remove' events by comparing two ports enumeration
// made at different times:
// - ports present in the new list but not in the old list are reported as 'added'
// - ports present in the old list but not in the new list are reported as 'removed'
func processUpdates(old, new []*enumerator.PortDetails, eventCB discovery.EventCallback) {
	for _, oldPort := range old {
		if !portListHas(new, oldPort) {
			if oldPort.VID == "1781" && oldPort.PID == "0941" {
				eventCB("remove", &discovery.Port{
					Address:  oldPort.Name,
					Protocol: "qnimble",
				})
			}
		}
	}

	for _, newPort := range new {
		if !portListHas(old, newPort) {
			if newPort.VID == "1781" && newPort.PID == "0941"  {
				eventCB("add", toDiscoveryPort(newPort))
			}
		}
	}
}

// portListHas checks if port is contained in list. The port metadata are
// compared in particular the port address, and vid/pid if the port is a usb port.
func portListHas(list []*enumerator.PortDetails, port *enumerator.PortDetails) bool {
	for _, p := range list {
		if port.Name == p.Name && port.IsUSB == p.IsUSB {
			if p.IsUSB &&
				port.VID == p.VID &&
				port.PID == p.PID &&
				port.MI == p.MI &&
				port.SerialNumber == p.SerialNumber {
				return true
			}
			if !p.IsUSB {
				return true
			}
		}
	}
	return false
}

func toDiscoveryPort(port *enumerator.PortDetails) *discovery.Port {
	props := properties.NewMap()
	if port.IsUSB {
		props.Set("vid", "0x"+port.VID)
		props.Set("pid", "0x"+port.PID)
		props.Set("mi", "0x"+port.MI)

		if port.VID == "1781" && port.PID == "0941" && port.MI == "00" {

			props.Set("upload", "1")
			props.Set("serialNumber", port.SerialNumber)

			res := &discovery.Port{
				Address:       port.Name,
				AddressLabel:  "Quarto (" + port.Name + ")",
				Protocol:      "qnimble",
				ProtocolLabel: "qnimble sam-ba emulator",
				Properties:    props,
			}
			return res
		} else {
			res := &discovery.Port{
				Address:       port.Name,
				AddressLabel:  port.Name,
				Protocol:      "serial",
				ProtocolLabel: "Serial Device",
				Properties:    props,
			}
			//return &discovery.Port{}
			return res

		}
	}
	return &discovery.Port{}
}
