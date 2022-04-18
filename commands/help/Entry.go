/**

Lux
Copyright (C) 2022  BanDev

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
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/api/keymanager"
	"github.com/fatih/color"
)

// Entry function is the entry point for
// the devices module.
func Entry() {

	// Get the latest commands from server
	var c Commands
	c.Fill()

	general.PrintHeading("Welcome to Lux!", color.FgWhite)
	general.PrintHeading("Lux " + general.BuildName, color.Italic)
	keymanager.PrintLuxHasAPIKey()
	general.Line()
	c.List()
	general.Line()
	general.PrintHeading("ABOUT LUX", color.FgWhite)
	general.PrintBoolParagraph("logged in", color.FgWhite, keymanager.LuxHasAPIKey())
	general.PrintStringParagraph("description", general.Description, color.FgWhite)
	general.PrintStringParagraph("build", general.BuildName + " (" + general.GetBits() + ")", color.FgWhite)
	general.PrintStringParagraph("license", general.License, color.FgWhite)
	general.PrintStringParagraph("repository", general.GHRepo, color.FgWhite)
	general.PrintBoolParagraph("up to date", color.FgWhite, general.BuildNo >= c.BuildNo)
}