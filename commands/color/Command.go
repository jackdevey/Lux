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
	"github.com/spf13/cobra"
)

// Command returns a pointer to the Color
// command, allowing the user to alter the
// color of their devices
func Command() *cobra.Command {
	// Set all flags
	Color.Flags().IntP("device", "d", 0, "the id of the Govee device")
	Color.Flags().BoolP("all", "a", false, "run the command on all devices")
	Color.Flags().StringP("hex", "x", "ffffff", "the hexadecimal color code to change the device to")
	Color.Flags().StringP("color", "c", "white", "the color to change the device to")
	// Either all or device flag must be present
	// but a check for this cannot be implemented
	// here and thus, must be implemented through
	// code.
	return Color
}
