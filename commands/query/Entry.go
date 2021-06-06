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

package query

import (
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/commands/devices"
	"github.com/fatih/color"
	"os"
	"strconv"
)

// Entry function is the entry point for
// the query module.
func Entry(args []string) {
	// Check a device id is present
	if len(args) <= 2 {
		general.PrintHeading("No device id provided. E.g. lux query 0", color.FgRed)
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
		general.PrintHeading("Device id provided is invalid", color.FgRed)
		return
	}

	// Find the device
	var d = ds.Data.Devices[dID]

	// Generate and parse the query
	var q Query
	q.Fill(c, d)

	// Print query to screen
	general.PrintHeading("QUERY RESULTS", color.FgWhite)
	general.PrintStringParagraph("Device:", q.Device, color.FgWhite)
	general.PrintStringParagraph("Model:", q.Model, color.FgWhite)
	general.PrintBoolParagraph("Online:", color.FgWhite, q.Online)
	general.PrintBoolParagraph("Power:", color.FgWhite, q.Power)
	general.PrintStringParagraph("Brightness:", strconv.Itoa(int(q.Brightness)) + "%", color.FgWhite)
	general.PrintStringParagraph("Colour:", q.Colour.ToString(), color.FgWhite)
}
