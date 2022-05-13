/*
 * Lux
 * Copyright (C) 2022 Jack Devey
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package color

import (
	"errors"
	"github.com/bandev/lux/api/core"
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/api/goveedevices"
	"github.com/bandev/lux/api/keymanager"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"regexp"
	"strconv"
)

// Color is the command that alters the color
// of the user's Govee devices
var Color = &cobra.Command{
	Use:   "color",
	Short: "change the color of a Govee device",
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
		// Create empty color object & hexStr
		var colorObj ControlCmdColor
		var hexStr string
		// If color shortcut provided, parse it here
		// or parse the color code
		if cmd.Flags().Changed("color") {
			// Get the corresponding color code
			var input, _ = cmd.Flags().GetString("color")
			var hex, err = shortcutColorLookup(input)
			if err != nil {
				return errors.New("unknown color value " + input)
			} else {
				hexStr = hex
			}
		} else {
			// Get hex code from input
			hexStr, _ = cmd.Flags().GetString("hex")
			var validColor, _ = regexp.MatchString("^([A-Fa-f0-9]{6})$", hexStr)
			if !validColor {
				return errors.New("color code should be in hexadecimal format, e.g. ffffff or 0067f4")
			}
		}
		// Parse color into format for Govee
		colorObj.R, _ = strconv.ParseInt(hexStr[0:2], 16, 64)
		colorObj.G, _ = strconv.ParseInt(hexStr[2:4], 16, 64)
		colorObj.B, _ = strconv.ParseInt(hexStr[4:6], 16, 64)
		// If all, then loop through all devices
		// otherwise do just singular
		var array []goveedevices.Device
		if all {
			array = core.GetAllDevices(connection)
		} else {
			array = []goveedevices.Device{core.GetDevice(device, connection)}
		}
		// Iterate through the provided devices
		for i, d := range array {
			// Create control and response structs
			// & send the data.
			var control Control
			var response Response
			response.Fill(control.Send(d, connection, colorObj))
			// Output data afterwards
			general.PrintHeading("Color "+strconv.Itoa(i), color.FgWhite)
			general.PrintStringParagraph("color", "#"+hexStr, color.FgWhite)
			general.PrintStringParagraph("transaction", response.Message, color.FgWhite)
		}
		return nil
	},
}
