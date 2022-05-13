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

package setup

import (
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/api/keymanager"
	"github.com/spf13/cobra"
)

// Setup guides the user through inputting
// their Govee API key, to link their account
// to the CLI.
var Setup = &cobra.Command{
	Use:   "setup",
	Short: "Link Lux to your Govee account",
	Run: func(cmd *cobra.Command, args []string) {
		// Is Lux already setup?
		if keymanager.HasKey() {
			general.PrintWarning("Lux is already setup with a Govee API Key")
			return
		}
		// Ask for API key if user has one
		// or show instructions to create one
		if HasAPIKey() {
			AskForKey()
		} else {
			ShowCreateKey()
		}
	},
}
