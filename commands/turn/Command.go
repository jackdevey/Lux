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

package turn

import "github.com/spf13/cobra"

// Command returns a pointer to Turn
// which turns a Govee device on or off
func Command() *cobra.Command {
	// Set all flags
	Turn.Flags().IntP("device", "d", 0, "the id of the Govee device")
	Turn.Flags().BoolP("all", "a", false, "run the command on all devices")
	Turn.Flags().BoolP("on", "t", false, "turn device on")
	Turn.Flags().BoolP("off", "f", false, "turn device off")
	return Turn
}
