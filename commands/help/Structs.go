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
	"encoding/json"
	"internal/design"
	"internal/general"
	"internal/structs"
	"pkg/color"
	"strings"
)

// Commands struct holds parsed json
// data about all the commands available
// on the CLI.
type Commands struct {
	Commands []Command `json:"commands"`
	BuildNo int `json:"build_no"`
}

// Command struct holds parsed json data
// about each specific command available
// on the CLI.
type Command struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Examples []string `json:"examples"`
	BuildNo int `json:"build_no"`
}

// Fill fills an empty struct with the latest
// commands from the commands api.
func (c *Commands) Fill() {
	var conn structs.Connection
	conn.Url = "https://bandev.uk"
	conn.Key = "NOT-NEEDED"
	_ = json.Unmarshal(conn.Get("/api/lux/commands.json"), &c)
}

// List lists all the commands to the screen
// in a similar way to devices.
func (c *Commands) List() {
	design.PrintHeading("COMMANDS", color.FgWhite)
	for i := 0; i < len(c.Commands); i++ {
		var command = c.Commands[i]
		if command.BuildNo != general.BuildNo { continue }
		design.PrintStringParagraph(command.Name, command.Description, color.FgWhite)
	}
	design.Line()
	design.PrintHeading("EXAMPLES", color.FgWhite)
	for i := 0; i < len(c.Commands); i++ {
		var command = c.Commands[i]
		if command.BuildNo != general.BuildNo { continue }
		design.PrintStringParagraph(command.Name, strings.Join(command.Examples, ", "), color.FgWhite)
	}
}