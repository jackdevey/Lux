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

package general

import (
	"github.com/fatih/color"
)

// PrintHeading prints text in a unified heading
// format.
func PrintHeading(s string, c color.Attribute) {
	d := color.New(c, color.Bold)
	_, _ = d.Println(s)
}

// PrintWarning prints a warning message to the screen
func PrintWarning(s string) {
	d := color.New(color.FgYellow)
	_, _ = d.Print("! ")
	d2 := color.New(color.FgWhite)
	_, _ = d2.Print(s)
	print("\r\n")
}

func PrintError(s string) {
	d := color.New(color.FgRed)
	_, _ = d.Print("✗ ")
	d2 := color.New(color.FgWhite)
	_, _ = d2.Print(s)
	print("\r\n")
}

// PrintNeutral prints a neutral message to the screen
func PrintNeutral(s string) {
	d := color.New(color.FgWhite)
	_, _ = d.Println("- " + s)
}

// PrintStringParagraph prints text in a unified paragraph
// format taking into account grid spacing.
func PrintStringParagraph(s1 string, s2 string, c1 color.Attribute) {
	d1 := color.New(c1)
	_, _ = d1.Print("  " + s1)
	_, _ = d1.Print(GridSpacing(s1) + s2)
	print("\r\n")
}

// PrintBoolParagraph is very specific function for
// printing in a format like X: [ YES / NO ]
func PrintBoolParagraph(s1 string, c1 color.Attribute, yes bool) {
	d1 := color.New(c1)
	_, _ = d1.Print("  " + s1)
	var c2 = color.FgRed
	if yes { c2 = color.FgGreen }
	d2 := color.New(c2)
	_, _ = d2.Print(GridSpacing(s1) + BoolToString(yes))
	print("\r\n")
}

// Line prints an empty line to the screen
func Line() {
	PrintHeading("", color.Concealed)
}

// GridSpacing calculates & prints spaces in order
// to make grid style print statements.
func GridSpacing(s1 string) string {
	var str = ""
	var spacing = 15 - len(s1)
	for i := 0; i <= spacing; i++ { str += " " }
	return str
}