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

package devices

import (
	"errors"
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/api/goveedevices"
	"github.com/bandev/lux/api/keymanager"
	"github.com/spf13/cobra"
)

// Devices is the command that returns
// a list of all devices in the user
// account.
var Devices = &cobra.Command{
	Use:   "devices",
	Short: "List devices on your Govee account",
	RunE: func(cmd *cobra.Command, args []string) error {
		// If lux is not setup, throw an error
		if !keymanager.PrintLuxHasAPIKey() {
			return errors.New("not setup with Govee API key")
		}
		// Setup connection
		var connection = general.NewGoveeConnection()
		// Get expert from flags
		var expert, _ = cmd.Flags().GetBool("expert")
		// Get limit from flags
		var limit, _ = cmd.Flags().GetInt("limit")
		// Get devices
		var devices goveedevices.Devices
		devices.Get(connection)
		// Decide on the command to run
		if expert {
			devices.ComplexList(limit)
		} else {
			devices.SimpleList(limit)
		}
		return nil
	},
}