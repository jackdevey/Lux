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

package turn

import (
	"commands/devices"
	"internal/general"
	"os"
	"pkg/color"
	"strconv"
	"strings"
)

// Entry point for the turn on commands.
// Checks the arguments and decides what
// is required to do.
func Entry(args []string) {
	// Check a device id is present
	if len(args) <= 2 {
		general.PrintHeading("No device id provided. E.g. lux turn 0 on", color.FgRed)
		return
	}else if len(args) <= 3 {
		general.PrintHeading("Not enough arguments required. E.g. lux turn 0 off", color.FgRed)
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
		general.PrintHeading("Device id provided is invalid", color.FgRed)
		return
	}

	// Find the command
	var cmd = general.StringToBool(args[3])

	// Find the device
	var d = ds.Data.Devices[dId]

	// Create control and response structs
	// & send the data.
	var control Control
	var response Response
	response.Fill(control.Send(d, c, cmd))

	// Output data afterwards
	general.PrintHeading("TURN 0 " + strings.ToUpper(args[2]), color.FgWhite)
	general.PrintBoolParagraph("power", color.FgWhite, cmd)
	general.PrintStringParagraph("transaction", response.Message, color.FgWhite)
}


