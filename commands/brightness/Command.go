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

package brightness

import (
	"github.com/spf13/cobra"
)

// Command returns a pointer to the Brightness
// command, allowing the user to alter the
// brightness of their devices
func Command() *cobra.Command {
	// Set all flags
	Brightness.Flags().IntP("device", "d", 0, "the id of the Govee device")
	Brightness.Flags().BoolP("all", "a", false, "run the command on all devices")
	Brightness.Flags().IntP("percent", "p", 50, "the percentage of brightness")
	// Either all or device flag must be present
	// but a check for this cannot be implemented
	// here and thus, must be implemented through
	// code.
	return Brightness
}
