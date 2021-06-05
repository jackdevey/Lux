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

package color

import (
	"commands/devices"
	"internal/general"
	"os"
	"pkg/color"
	"strconv"
	"strings"
)

// Entry point for the colour on commands.
// Checks the arguments and decides what
// is required to do.
func Entry(args []string) {
	// Check a device id is present
	if len(args) <= 2 {
		general.PrintHeading("No device id provided. E.g. lux color 0 #0067f4", color.FgRed)
		return
	}else if len(args) <= 3 {
		general.PrintHeading("Not enough arguments required. E.g. lux color 0 #0067f4", color.FgRed)
		return
	}

	// Determine the device id from the args
	var dId, _ = strconv.Atoi(args[2])

	// Create a new connection struct
	var c general.Connection
	c.Key = os.Getenv("GOVEE_API_KEY")
	c.Url = "https://developer-api.govee.com/"

	// Get a list of devices owned by the user
	var ds devices.Devices
	ds.Get(c)

	// Check dId provided is valid
	if len(ds.Data.Devices) <= dId || dId < 0 {
		general.PrintHeading("Device id provided is invalid. E.g. lux color 0 #0067f4", color.FgRed)
		return
	}

	// Check for hexadecimal colour codes
	var hex = args[3]
	var colour ControlCmdColor
	if strings.HasPrefix(hex, "#") && len(hex) == 7 {
		// Parse hexadecimal colours into rgb
		// values
		colour.R, _ = strconv.ParseInt(hex[1:3], 16, 64)
		colour.G, _ = strconv.ParseInt(hex[3:5], 16, 64)
		colour.B, _ = strconv.ParseInt(hex[5:7], 16, 64)
	}else {
		general.PrintHeading("No hexadecimal colour code provided. E.g. lux color 0 #0067f4", color.FgRed)
		return
	}


	// Find the device
	var d = ds.Data.Devices[dId]

	// Create control and response structs
	// & send the data.
	var control Control
	var response Response
	response.Fill(control.Send(d, c, colour))

	// Output data afterwards
	general.PrintHeading("COLOR " + args[2] + " " + strings.ToUpper(args[3]), color.FgWhite)
	general.PrintStringParagraph("colour", args[3], color.FgWhite)
	general.PrintStringParagraph("transaction", response.Message, color.FgWhite)
}


