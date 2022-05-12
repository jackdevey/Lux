/**

Lux
Copyright (C) 2022  Jack Devey

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
	"github.com/bandev/lux/commands/color"
	"github.com/bandev/lux/commands/misc"
	"github.com/bandev/lux/commands/setup"
	"github.com/spf13/cobra"
)

// main is the command that is first
// run. It determines where to send the
// user based on what the arguments are.
func main() {
	// Decide what part of the cli to run

	var root = &cobra.Command{
		Use: "lux",
	}

	root.AddCommand(misc.Version)
	root.AddCommand(setup.Setup)
	root.AddCommand(color.Command())

	root.Execute()
}
