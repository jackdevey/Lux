/**

Lux
Copyright (C) 2022  Jack Devey

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

*/

package core

import (
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/api/goveedevices"
	"github.com/fatih/color"
	"os"
	"strconv"
)

// GetDevicesFrom will return an array of devices
// from Govee
func GetDevicesFrom(arg string, c general.Connection) []goveedevices.Device {
	var ds goveedevices.Devices
	ds.Get(&c)
	if arg == "@a" {
		return ds.Data.Devices
	}
	var dID, _ = strconv.Atoi(arg)
	if len(ds.Data.Devices) <= dID || dID < 0 {
		general.PrintHeading("Device id provided is invalid", color.FgRed)
		os.Exit(1)
	}
	return []goveedevices.Device{ds.Data.Devices[dID]}
}

func GetAllDevices(c *general.Connection) []goveedevices.Device {
	var ds goveedevices.Devices
	ds.Get(c)
	return ds.Data.Devices
}

func GetDevice(id int, c *general.Connection) goveedevices.Device {
	var ds goveedevices.Devices
	ds.Get(c)
	return ds.Data.Devices[id]
}
