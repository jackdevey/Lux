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
	"flag"
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/api/keymanager"
	"os"
)

// Entry function is the entry point for
// the devices module.
func Entry(args []string) {
	// If lux is not setup, throw an error
	if !keymanager.PrintLuxHasAPIKey() { return }

	// Use flags package to parse flags
	// provided to the sub command
	devicesCmd := flag.NewFlagSet("devices",flag.ExitOnError)
	expert := devicesCmd.Bool("expert", false, "show more technical device details")
	limit := devicesCmd.Int("limit", 0, "add a limit to the number of devices shown")

	// Parse flags & handle errors
	err := devicesCmd.Parse(os.Args[2:])
	if err != nil {
		return
	}

	// Create a new connection struct
	var c general.Connection
	c.Key = keymanager.GetAPIKey()
	c.Base = "https://developer-api.govee.com/"

	// Get a list of devices from Govee
	var devices Devices
	devices.Get(c)

	// Decide on the command to run
	if *expert {
		devices.ComplexList(*limit)
	} else {
		devices.SimpleList(*limit)
	}
}