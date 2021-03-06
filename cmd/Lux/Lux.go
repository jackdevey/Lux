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
	"github.com/bandev/lux/commands/brightness"
	"github.com/bandev/lux/commands/color"
	"github.com/bandev/lux/commands/devices"
	"github.com/bandev/lux/commands/misc"
	"github.com/bandev/lux/commands/query"
	"github.com/bandev/lux/commands/setup"
	"github.com/bandev/lux/commands/turn"
	"github.com/spf13/cobra"
)

// main is the command that is first
// run. It determines where to send the
// user based on what the arguments are.

// Decide what part of the cli to run
var root = &cobra.Command{
	Use: "lux",
	Short: "Lux is a command-line interface for controlling and monitoring Govee lighting",
}

func main() {
	// Add all commands
	root.AddCommand(color.Command())
	root.AddCommand(devices.Command())
	root.AddCommand(turn.Command())
	root.AddCommand(brightness.Command())
	root.AddCommand(query.Command())
	root.AddCommand(misc.Version)
	root.AddCommand(setup.Setup)
	// Execute the command
	_ = root.Execute()
}
