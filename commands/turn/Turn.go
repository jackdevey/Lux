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

package turn

import (
	"errors"
	"github.com/bandev/lux/api/core"
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/api/goveedevices"
	"github.com/bandev/lux/api/keymanager"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strconv"
)

// Turn is the command that allows
// the user to turn a device on or off
var Turn = &cobra.Command{
	Use:   "turn",
	Short: "Turn a device on or off",
	RunE: func(cmd *cobra.Command, args []string) error {
		// If lux is not setup, throw an error
		if !keymanager.PrintLuxHasAPIKey() {
			return errors.New("not setup with Govee API key")
		}
		// Setup connection
		var connection = general.NewGoveeConnection()
		// Get all from flags
		var all, _ = cmd.Flags().GetBool("all")
		// Get device from flags
		var device, _ = cmd.Flags().GetInt("device")
		if !cmd.Flags().Changed("device") && !all {
			return errors.New("no device id provided")
		}
		// Work out power choice
		// Get all from flags
		if !cmd.Flags().Changed("on") && !cmd.Flags().Changed("off") {
			return errors.New("no power option provided")
		}
		// if not on then must be off
		var option, _ = cmd.Flags().GetBool("on")
		// If all, then loop through all devices
		// otherwise do just singular
		var array []goveedevices.Device
		if all {
			array = core.GetAllDevices(connection)
		} else {
			dev, err := core.GetDevice(device, connection)
			// If err, pass it on
			if err != nil { return err }
			array = []goveedevices.Device{dev}
		}
		// Iterate through the provided devices
		for i, d := range array {
			// Create control and response structs
			// & send the data.
			var control Control
			var response Response
			response.Fill(control.Send(d, connection, option))
			// Output data afterwards
			general.PrintHeading("Turn " + strconv.Itoa(i), color.FgWhite)
			general.PrintBoolParagraph("power", color.FgWhite, option)
			general.PrintStringParagraph("transaction", response.Message, color.FgWhite)
		}
		return nil
	},
}


