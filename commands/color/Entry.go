/**

Lux
Copyright (C) 2022  BanDev

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
	"github.com/bandev/lux/api/core"
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/api/keymanager"
	"github.com/fatih/color"
	"regexp"
	"strconv"
	"strings"
)

// Entry point for the colour on commands.
// Checks the arguments and decides what
// is required to do.
func Entry(args []string) {
	// If lux is not setup, throw an error
	if !keymanager.PrintLuxHasAPIKey() { return }

	// Check a device id is present
	if len(args) <= 2 {
		general.PrintHeading("No device id provided. E.g. lux color 0 0067f4", color.FgRed)
		return
	}else if len(args) <= 3 {
		general.PrintHeading("Not enough arguments required. E.g. lux color 0 0067f4", color.FgRed)
		return
	}

	// Create a new connection struct
	var c general.Connection
	c.Key = keymanager.GetAPIKey()
	c.Base = "https://developer-api.govee.com/"

	// Check for hexadecimal colour codes
	var hex = args[3]
	var colour ControlCmdColor

	// Remove # if it was added
	hex = strings.ReplaceAll(hex, "#", "")

	// Check the colour code is valid
	var validHex, _ = regexp.MatchString("^([A-Fa-f0-9]{6})$", hex)

	if validHex {
		// Parse hexadecimal colours into rgb
		// values
		colour.R, _ = strconv.ParseInt(hex[0:2], 16, 64)
		colour.G, _ = strconv.ParseInt(hex[2:4], 16, 64)
		colour.B, _ = strconv.ParseInt(hex[4:6], 16, 64)
	}else {
		general.PrintHeading(hex + " is not a valid colour code. E.g. lux color 0 0067f4", color.FgRed)
		return
	}

	// For each device
	for i, d := range core.GetDevicesFrom(args[2], c) {
		// Create control and response structs
		// & send the data.
		var control Control
		var response Response
		response.Fill(control.Send(d, c, colour))

		// Output data afterwards
		general.PrintHeading("COLOR " + strconv.Itoa(i) + " #" + strings.ToUpper(hex), color.FgWhite)
		general.PrintStringParagraph("colour", "#" + hex, color.FgWhite)
		general.PrintStringParagraph("transaction", response.Message, color.FgWhite)
	}
}


