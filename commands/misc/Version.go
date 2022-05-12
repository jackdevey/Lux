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

package misc

import (
	"fmt"
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/api/github"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Version prints the current lux version
// and checks for any available updates using
// the GitHub REST API
var Version = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Lux",
	Run: func(cmd *cobra.Command, args []string) {

		var upgrade = false
		// Get the latest GitHub release for lux
		var latest = github.GetLatestRelease()
		upgrade = latest.Tag != general.BuildName
		// If an upgrade is available, print a notice
		if upgrade {
			general.PrintWarning("Lux " + latest.Tag + " is out. Upgrade now")
			fmt.Println()
		}
		// Print the lux version details
		general.PrintHeading("Lux version", color.FgWhite)
		general.PrintStringParagraph("tag", general.BuildName, color.FgWhite)
		general.PrintStringParagraph("architecture", general.GetBits(), color.FgWhite)
		general.PrintBoolParagraph("up to date", color.FgWhite, !upgrade)

	},
}
