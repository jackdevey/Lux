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

package general

// BoolToString converts a Boolean to a
// string. Makes boolean values more readable
// to a user. When true outputs Yes.
func BoolToString(b bool) string {
	if b {
		return "[✓ Yes]"
	}else {
		return "[✗ No]"
	}
}

// StringToBool converts a string to a
// boolean. Sorta like the reverse of
// BoolToString. It will convert words like
// on / off, true / false into boolean data.
// Anything else will return false.
func StringToBool(s string) bool {
	if s == "true" || s == "on" || s == "yes" {
		return true
	}else {
		return false
	}
}

// HexTo2Places is used for converting
// the value 0 to 00. This is used for
// hexadecimal colour codes.
func HexTo2Places(h string) string {
	if h == "0" {
		return "00"
	}else {
		return h
	}
}
