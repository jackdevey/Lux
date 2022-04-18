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
	"encoding/json"
	"github.com/bandev/lux/api/general"
	"github.com/fatih/color"
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
	var conn general.Connection
	conn.Base = "https://bandev.uk"
	conn.Key = "NOT-NEEDED"
	_ = json.Unmarshal(conn.Get("/api/lux/commands.json"), &c)
}

// List lists all the commands to the screen
// in a similar way to devices.
func (c *Commands) List() {
	general.PrintHeading("COMMANDS", color.FgWhite)
	for i := 0; i < len(c.Commands); i++ {
		var command = c.Commands[i]
		if command.BuildNo != general.BuildNo { continue }
		general.PrintStringParagraph(command.Name, command.Description, color.FgWhite)
	}
	general.Line()
	general.PrintHeading("EXAMPLES", color.FgWhite)
	for i := 0; i < len(c.Commands); i++ {
		var command = c.Commands[i]
		if command.BuildNo != general.BuildNo { continue }
		general.PrintStringParagraph(command.Name, strings.Join(command.Examples, ", "), color.FgWhite)
	}
}