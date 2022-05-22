/**

Lux
Copyright (C) 2022  Jack Devey

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

package keymanager

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ApiKey stores the user's
// API key.
type ApiKey struct {
	Key string `json:"key"`
}

// Store the current key to
// the user's storage
func (k *ApiKey) Store() {
	dir, _ := os.UserHomeDir()
	path := dir + "\\.config\\lux"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.MkdirAll(path, os.ModePerm)
	}
	var f *os.File
	if _, err := os.Stat(path + "\\key.json"); os.IsNotExist(err) {
		f, _ = os.Create(path + "\\key.json")
	} else {
		f, _ = os.OpenFile(path+"\\key.json", os.O_CREATE|os.O_WRONLY, 0644)
	}
	b, _ := json.Marshal(&k)
	_, e := f.Write(b)
	if e != nil {
		print(e.Error())
	}
	_ = f.Sync()
}

// Extract the key from the user's
// storage
func (k *ApiKey) Extract() {
	dir, _ := os.UserHomeDir()
	path := dir + "\\.config\\lux\\key.json"
	b, _ := ioutil.ReadFile(path)
	_ = json.Unmarshal(b, &k)
}

// HasKey check if the software is setup
func HasKey() bool {
	return GetAPIKey() != ""

}

// PrintLuxHasAPIKey uses LuxHasAPIKey to
// print an error if the user has no api key.
// DOES NOT PRINT IS BUG NEEDS FIX
func PrintLuxHasAPIKey() bool {
	if !HasKey() {
		return false
	}
	return true
}

// GetAPIKey allows packages to
// request the api key
func GetAPIKey() string {
	var key ApiKey
	key.Extract()
	return key.Key
}
