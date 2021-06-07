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

package setup

import (
	"fmt"
	"github.com/bandev/lux/api/general"
	"github.com/cli/browser"
)

// Entry command for the setup function,
// asks the user if they have an API Key,
// and if they don't prompts them to a tutorial
// on how to obtain one.
func Entry() {
	general.PrintWarning("To use Lux you must have a Govee API Key")

	if HasAPIKey() {
		GetAPIKey()
	}else {
		general.PrintNeutral("Press enter to open Api Key creation guide in browser...")
		_, _ = fmt.Scanln()
		_ = browser.OpenURL("https://github.com/bandev/lux")
	}
}
