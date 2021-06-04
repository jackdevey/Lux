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

package devices

import (
	"encoding/json"
	"internal/general"
	"pkg/color"
	"strconv"
	"strings"
)

// Devices struct maps to the Govee Api
// JSON Endpoints as of v1. Devices
// is a list of all the Device's that
// belong to an authenticated user.
type Devices struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	Status  int    `json:"status"`
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
	Commands []string `json:"supportCmds"`
}

// Get Fills an empty Devices object with information
// provided from Govee's API using the Connection
// that must be passed as a parameter.
func (d *Devices) Get(c general.Connection) {
	err := json.Unmarshal(c.Get("v1/devices"), &d)
	if err != nil { println("Error while parsing devices object") }
}

// SimpleList prints a simple list of Device's to
// the terminal, the list includes the name,
// model and if the device is controllable or
// not by this app.
func (d *Devices) SimpleList() {
	for i := 0; i < len(d.Data.Devices); i++ {
		var device = d.Data.Devices[i]
		general.PrintHeading("DEVICE " + strconv.Itoa(i), color.FgWhite)
		general.PrintStringParagraph("Name:", device.Name, color.FgWhite)
		general.PrintStringParagraph("Model:", device.Model, color.FgWhite)
		general.PrintBoolParagraph("Controllable: ", color.FgWhite, device.Controllable)
	}
}

// ComplexList outputs the User's devices in a
// nice but complex way. It pretty much dumps
// everything the app knows about their device
// to the console.
func (d *Devices) ComplexList() {
	for i := 0; i < len(d.Data.Devices); i++ {
		var device = d.Data.Devices[i]
		general.PrintHeading("DEVICE " + strconv.Itoa(i), color.FgWhite)
		general.PrintStringParagraph("MAC Address:", device.MAC, color.FgWhite)
		general.PrintStringParagraph("Model:", device.Model, color.FgWhite)
		general.PrintStringParagraph("Name:", device.Name, color.FgWhite)

		general.PrintBoolParagraph("Controllable: ", color.FgWhite, device.Controllable)
		general.PrintBoolParagraph("Retrievable: ", color.FgWhite, device.Controllable)

		general.PrintStringParagraph("Commands:", strings.Join(device.Commands, ", "), color.FgWhite)
	}
}