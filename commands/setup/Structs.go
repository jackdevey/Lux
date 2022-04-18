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

package setup

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/api/keymanager"
)

// HasAPIKey is used to check if the user has
// an api key.
func HasAPIKey() bool {
	var question = []*survey.Question{
		{
			Name: "HasAPIKey",
			Prompt: &survey.Select{
				Message: "Do you have one?",
				Options: []string{"yes", "no"},
				Default: "yes",
			},
		},
	}
	var answer = struct {
		HasAPIKey string
	}{}
	_ = survey.Ask(question, &answer)
	return general.StringToBool(answer.HasAPIKey)
}

// GetAPIKey is used to read the user's
// api key.
func GetAPIKey() string {
	var question = []*survey.Question{
		{
			Name: "GetAPIKey",
			Prompt: &survey.Input{Message: "Enter Govee Api Key"},
			Validate: survey.Required,
			Transform: survey.Title,
		},
	}
	var answer = struct {
		GetAPIKey string
	}{}
	_ = survey.Ask(question, &answer)
	return answer.GetAPIKey
}

// AskForKey uses the GetAPIKey function
// to ask the user for and validate their
// api key.
func AskForKey() {
	key := GetAPIKey()
	c := general.Connection{
		Key: key,
		Base: "https://developer-api.govee.com/",
	}
	if c.TestKey() {
		keymngr := keymanager.KeyStore{Key: key}
		keymngr.Store()
	}else if key != "" {
		general.PrintError("Invalid API Key")
		AskForKey()
	}
}

