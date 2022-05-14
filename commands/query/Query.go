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

package query

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

// Query is the command that allows
// the user to view the status of a
// device
var Query = &cobra.Command{
	Use: "query",
	Short: "Query the state of a device",
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
			// Generate and parse the query
			var q DeviceQuery
			q.Fill(connection, d)
			// Print query to screen
			general.PrintHeading("Query "+strconv.Itoa(i), color.FgWhite)
			general.PrintStringParagraph("device", q.Device, color.FgWhite)
			general.PrintStringParagraph("model", q.Model, color.FgWhite)
			general.PrintBoolParagraph("online", color.FgWhite, q.Online)
			general.PrintBoolParagraph("power", color.FgWhite, q.Power)
			general.PrintStringParagraph("brightness", strconv.Itoa(int(q.Brightness))+"%", color.FgWhite)
			general.PrintStringParagraph("color", q.Colour.ToString(), color.FgWhite)
		}
		return nil
	},
}