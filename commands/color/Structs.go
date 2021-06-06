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

package color

import (
	"encoding/json"
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/commands/devices"
)

// Control struct manages the control
// json data that will be send to govee.
type Control struct {
	Device string `json:"device"`
	Model string `json:"model"`
	Cmd ControlCmd `json:"cmd"`
}

// ControlCmd is the control command
// json that will be sent to govee.
type ControlCmd struct {
	Name string `json:"name"`
	Value ControlCmdColor `json:"value"`
}

// ControlCmdColor is the struct for
// managing the RGB values for each
// colour.
type ControlCmdColor struct {
	R int64 `json:"r"`
	G int64 `json:"g"`
	B int64 `json:"b"`
}

// Send sends all the data to govee to set
// the value of the colour
func (c *Control) Send (d devices.Device, conn general.Connection, color ControlCmdColor) []byte {
	c.Device = d.MAC
	c.Model = d.Model
	c.Cmd.Name = "color"

	c.Cmd.Value = color

	var data, _ = json.Marshal(&c)
	return conn.Put("v1/devices/control", data)
}

// Response struct is for parsing the json
// returned by govee.
type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

// Fill fills the response struct with the
// data provided by govee.
func (r *Response) Fill (c []byte) {
	_ = json.Unmarshal(c, &r)
}