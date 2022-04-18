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

package main

import (
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/commands/brightness"
	"github.com/bandev/lux/commands/color"
	"github.com/bandev/lux/commands/devices"
	"github.com/bandev/lux/commands/help"
	"github.com/bandev/lux/commands/query"
	"github.com/bandev/lux/commands/setup"
	"github.com/bandev/lux/commands/turn"
	"os"
)

// main is the command that is first
// run. It determines where to send the
// user based on what the arguments are.
func main() {
	// Decide what part of the cli to run

	// If no argument given, run help
	if len(os.Args) == 1 { help.Entry(); return }

	// If an arg is present check what was
	// requested
	switch os.Args[1] {
		case "devices": devices.Entry(os.Args)
		case "query": query.Entry(os.Args)
		case "turn": turn.Entry(os.Args)
		case "brightness": brightness.Entry(os.Args)
		case "color": color.Entry(os.Args)
		case "setup": setup.Entry()
		case "help": help.Entry()
		default: general.PrintError("Unknown command " + os.Args[1])
	}
}

