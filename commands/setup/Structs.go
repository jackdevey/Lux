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
	"github.com/AlecAivazis/survey/v2"
	"github.com/bandev/lux/api/general"
)

// hasApiKey is used to check if the user has
// an api key.
func hasApiKey() bool {
	var question = []*survey.Question{
		{
			Name: "HasApiKey",
			Prompt: &survey.Select{
				Message: "Do you have one?",
				Options: []string{"yes", "no"},
				Default: "yes",
			},
		},
	}
	var answer = struct {
		HasApiKey string
	}{}
	_ = survey.Ask(question, &answer)
	return general.StringToBool(answer.HasApiKey)
}

// getApiKey is used to read the user's
// api key.
func getApiKey() string {
	var question = []*survey.Question{
		{
			Name: "GetApiKey",
			Prompt: &survey.Input{Message: "Enter Govee Api Key"},
			Validate: survey.Required,
			Transform: survey.Title,
		},
	}
	var answer = struct {
		GetApiKey string
	}{}
	_ = survey.Ask(question, &answer)
	return answer.GetApiKey
}
