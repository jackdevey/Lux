/*
 * Lux
 * Copyright (C) 2022 Jack Devey
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package color

import "errors"

func shortcutColorLookup(input string) (string, error) {
	switch input {
	// Core colors
	case "white": return "ffffff", nil
	case "red": return "ff0000", nil
	case "green": return "00ff00", nil
	case "blue": return "0000ff", nil
	// Other colors
	case "pink": return "ff69b4", nil
	case "orange": return "ffa500", nil
	case "yellow": return "ffff00", nil
	case "purple": return "800080", nil
	case "teal": return "008080", nil
	case "brown": return "8b4513", nil
	case "gray": return "808080", nil
	default: return "", errors.New("unknown input")
	}
}
