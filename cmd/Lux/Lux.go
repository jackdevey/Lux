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
	"internal/design"
	"internal/structs"
	"os"
	"pkg/fatih/color"
)

var connection structs.Connection

// This is temporary while testing things
func main() {

	connection.Key = os.Getenv("GOVEE_API_KEY")
	connection.Url = "https://developer-api.govee.com/"

	switch os.Args[1] {
		case "devices":
			Devices()
			break
	}

}

// Devices checks the second arguments given
// and then returns a list of either simple
// or complex devices.
func Devices() {
	var command string
	if len(os.Args) > 2 { command = os.Args[2] } else { command = "simple" }
	var devices structs.Devices
	devices.Get(connection)
	if command == "complex" {
		devices.ComplexList()
	}else if command == "simple" {
		devices.SimpleList()
	}else {
		design.PrintHeading("Unknown command '" + command + "' for devices", color.FgRed)
	}
}

