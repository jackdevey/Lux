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

package structs

import (
	"internal/general"
	"strconv"
)

// Colour defines the colour values
// used for interfacing with Govee's
// api.
type Colour struct {
	R int64
	G int64
	B int64
}

// ToString converts a Colour struct
// into a hexadecimal colour code.
func (c *Colour) ToString() string {
	var hr = general.HexTo2Places(strconv.FormatInt(c.R, 16))
	var hg = general.HexTo2Places(strconv.FormatInt(c.G, 16))
	var hb = general.HexTo2Places(strconv.FormatInt(c.B, 16))
	return "#" + hr + hg + hb
}