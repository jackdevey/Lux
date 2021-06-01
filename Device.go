/**

Lux
Copyright (C) 2021  BanDev

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

package main

import (
	"encoding/json"
	"strconv"
)

// Devices struct maps to the Govee Api
// JSON Endpoints as of v1. Devices
// is a list of all the Device's that
// belong to an authenticated user.
type Devices struct {
	Data Data `json:"data"`
	Message string `json:"message"`
	Status int `json:"status"`
}

// Data struct is so annoying but necessary
// to use the json.Unmarshal func provided
// by Go Lang because of the API's use of
// data padding. I wish I could delete it,
// but I dont think I can :(
type Data struct {
	Devices []Device `json:"devices"`
}

// Device struct is a map of the API's device
// node. The struct contains information about
// the user's device.
type Device struct {
	MAC string `json:"device"`
	Model string `json:"model"`
	Name string `json:"deviceName"`
	Controllable bool `json:"controllable"`
	Retrievable bool `json:"retrievable"`
	Commands [4]string `json:"supportCmds"`
}

// Fills an empty Devices object with information
// provided from Govee's API using the Connection
// that must be passed as a parameter.
func (d *Devices) get(c Connection) {
	err := json.Unmarshal(c.get("v1/devices"), &d)
	if err != nil { println("Error while parsing devices object") }
}

// Prints a simple list of Device's to the terminal
// the list includes the name, model and if the
// device is controllable or not by this app.
func (d *Devices) simpleList() {
	for i := 0; i < len(d.Data.Devices); i++ {
		var device = d.Data.Devices[i]
		println(strconv.Itoa(i) + " {")
		println("   Name: " + device.Name)
		println("   Model: " + device.Model)
		println("   Controllable: " + boolToString(device.Controllable))
		println("}")
	}
}