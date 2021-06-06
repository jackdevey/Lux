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

package brightness

import (
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/commands/devices"
	"github.com/fatih/color"
	"os"
	"strconv"
)

// Entry point for the brightness on commands.
// Checks the arguments and decides what
// is required to do.
func Entry(args []string) {
	// Check a device id is present
	if len(args) <= 2 {
		general.PrintHeading("No device id provided. E.g. lux brightness 0 100", color.FgRed)
		return
	}else if len(args) <= 3 {
		general.PrintHeading("Not enough arguments required. E.g. lux brightness 0 69", color.FgRed)
		return
	}

	// Determine the device id from the args
	var dID, _ = strconv.Atoi(args[2])

	// Create a new connection struct
	var c general.Connection
	c.Key = os.Getenv("GOVEE_API_KEY")
	c.Base = "https://developer-api.govee.com/"

	// Get a list of devices owned by the user
	var ds devices.Devices
	ds.Get(c)

	// Check dId provided is valid
	if len(ds.Data.Devices) <= dID || dID < 0 {
		general.PrintHeading("Device id provided is invalid. E.g. lux brightness 0 50", color.FgRed)
		return
	}

	// Find the power
	var power, err = strconv.Atoi(args[3])
	if err != nil {
		general.PrintHeading("Brightness must be a number", color.FgRed)
		return
	}

	// Check the power is allowed
	if 0 > power || power > 100 {
		general.PrintHeading("Brightness must be in range 0 < brightness < 100", color.FgRed)
		return
	}

	// Find the device
	var d = ds.Data.Devices[dID]

	// Create control and response structs
	// & send the data.
	var control Control
	var response Response
	response.Fill(control.Send(d, c, power))

	// Output data afterwards
	general.PrintHeading("BRIGHTNESS " + args[2] + " " + args[3] + "%", color.FgWhite)
	general.PrintStringParagraph("brightness", strconv.Itoa(power) + "%", color.FgWhite)
	general.PrintStringParagraph("transaction", response.Message, color.FgWhite)
}


