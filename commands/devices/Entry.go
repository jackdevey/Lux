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
	"github.com/bandev/lux/api/general"
	"github.com/fatih/color"
	"os"
)

// Entry function is the entry point for
// the devices module.
func Entry(args []string) {
	// Determine the command the user wants
	var command string
	if len(args) > 2 { command = args[2] } else { command = "simple" }

	// Create a new connection struct
	var c general.Connection
	c.Key = os.Getenv("GOVEE_API_KEY")
	c.Base = "https://developer-api.govee.com/"

	// Get a list of devices from Govee
	var devices Devices
	devices.Get(c)

	// Device on the command to run
	if command == "complex" {
		devices.ComplexList()
	}else if command == "simple" {
		devices.SimpleList()
	}else {
		general.PrintHeading("Unknown command '" + command + "' for devices", color.FgRed)
	}
}