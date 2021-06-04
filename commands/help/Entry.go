/**

Lux
Copyright (C) 2021  BanDev

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

package help

import (
	"internal/design"
	"internal/general"
	"pkg/color"
)

// Entry function is the entry point for
// the devices module.
func Entry(args []string) {

	// Get latest commands from server
	var c Commands
	c.Fill()

	design.PrintHeading("Welcome to Lux!", color.FgWhite)
	design.PrintHeading("Lux " + general.BuildName, color.Italic)
	design.Line()
	c.List()
	design.Line()
	design.PrintHeading("ABOUT LUX", color.FgWhite)
	design.PrintStringParagraph("description", general.Description, color.FgWhite)
	design.PrintStringParagraph("build", general.BuildName, color.FgWhite)
	design.PrintStringParagraph("license", general.License, color.FgWhite)
	design.PrintStringParagraph("repository", general.GHRepo, color.FgWhite)
	design.PrintBoolParagraph("up to date", color.FgWhite, general.BuildNo == c.BuildNo)
}
