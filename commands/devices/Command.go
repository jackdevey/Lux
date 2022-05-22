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

package devices

import "github.com/spf13/cobra"

// Command returns a pointer to Devices
// which returns a list of all Govee
// devices registered to the user's
// account
func Command() *cobra.Command {
	// Set all flags
	Devices.Flags().BoolP("expert", "e", false, "show more technical device details")
	Devices.Flags().IntP("limit", "p", 0, "add a limit to the number of devices shown")
	return Devices
}