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

package query

import (
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/commands/devices"
	"github.com/buger/jsonparser"
)

// Query is the struct for storing the
// data returned by Govee.
type Query struct {
	Device     string
	Model      string
	Online     bool
	Power      bool
	Brightness int64
	Colour     general.Colour
}

// Fill fills a Query struct with
// the data returned from Govee with
// the parameter of a Device struct.
func (q *Query) Fill(c general.Connection, d devices.Device) {
	var r = c.Get("v1/devices/state?device=" + d.MAC + "&model=" + d.Model)
	q.Device, _ = jsonparser.GetString(r, "data", "device")
	q.Model, _ = jsonparser.GetString(r, "data", "model")

	var o, _ = jsonparser.GetString(r, "data", "properties", "[0]", "online")
	q.Online = general.StringToBool(o)

	var p, _ = jsonparser.GetString(r, "data", "properties", "[1]", "powerState")
	q.Power = general.StringToBool(p)

	q.Brightness, _ = jsonparser.GetInt(r, "data", "properties", "[2]", "brightness")

	q.Colour.R, _ = jsonparser.GetInt(r, "data", "properties", "[3]", "color", "r")
	q.Colour.G, _ = jsonparser.GetInt(r, "data", "properties", "[3]", "color", "g")
	q.Colour.B, _ = jsonparser.GetInt(r, "data", "properties", "[3]", "color", "b")
}