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

package main

import (
	"commands/devices"
	"commands/help"
	"commands/query"
	"os"
)

// main is the command that is first
// run. It determines where to send the
// user based on what the arguments are.
func main() {
	// Decide what part of the cli to run

	// If no argument given, run help
	if len(os.Args) == 1 { help.Entry(os.Args); return }

	// If an arg is present check what was
	// requested
	switch os.Args[1] {
		case "devices": devices.Entry(os.Args); break
		case "query": query.Entry(os.Args); break
		case "help": help.Entry(os.Args); break
	}
}

