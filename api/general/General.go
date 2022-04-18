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

package general

// BoolToString converts a Boolean to a
// string. Makes boolean values more readable
// to a user. When true outputs Yes.
func BoolToString(b bool) string {
	if b {
		return "[✓ Yes]"
	}
	return "[✗ No]"
}

// StringToBool converts a string to a
// boolean. Sorta like the reverse of
// BoolToString. It will convert words like
// on / off, true / false into boolean data.
// Anything else will return false.
func StringToBool(s string) bool {
	if s == "true" || s == "on" || s == "yes" {
		return true
	}
	return false
}

// HexTo2Places is used for converting
// the value 0 to 00. This is used for
// hexadecimal colour codes.
func HexTo2Places(h string) string {
	if h == "0" {
		return "00"
	}
	return h
}

// Below are a few constants that define info
// about Lux

// BuildName is the name of the build
const BuildName = "v1.1.1"
// BuildNo is the build number
const BuildNo = 3
// GHRepo is the repo url
const GHRepo = "https://github.com/bandev/lux"
// License is the license
const License = "GPL-3.0"
// Description is the description
const Description = "Lux is a cli for controlling and monitoring Govee lighting strips"

// GetBits returns x64 or x86 depending on the
// installed version
func GetBits() string {
	if uint64(^uintptr(0)) == ^uint64(0) {
		return "x64"
	}
	return "x86"
}

