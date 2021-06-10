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

package keymanager

import (
	"encoding/json"
	"github.com/bandev/lux/api/general"
	"io/ioutil"
	"os"
)

// KeyStore manages the user's
// API key
type KeyStore struct {
	Key string `json:"key"`
}

// Store the current key to
// the user's storage
func (k *KeyStore) Store() {
	dir, _ := os.UserHomeDir()
	path := dir + "\\.config\\lux"
	if _, err := os.Stat(path); os.IsNotExist(err){
		_ = os.MkdirAll(path, os.ModePerm)
	}
	var f *os.File
	if _, err := os.Stat(path+"\\key.json"); os.IsNotExist(err){
		f, _ = os.Create(path+"\\key.json")
	}else {
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
func (k *KeyStore) Extract() {
	dir, _ := os.UserHomeDir()
	path := dir + "\\.config\\lux\\key.json"
	b, _ := ioutil.ReadFile(path)
	_ = json.Unmarshal(b, &k)
}

// LuxHasAPIKey uses the KeyStore class to
// check if the software is setup
func LuxHasAPIKey() bool {
	var keymngr KeyStore
	keymngr.Extract()
	return keymngr.Key != ""
}

// PrintLuxHasAPIKey uses LuxHasAPIKey to
// print an error if the user has no api key.
func PrintLuxHasAPIKey() bool {
	if !LuxHasAPIKey() {
		general.PrintError("Lux is not setup! Please run lux setup to add your API Key")
		return false
	}
	return true
}

// GetAPIKey allows packages to
// request the api key
func GetAPIKey() string {
	var keymngr KeyStore
	keymngr.Extract()
	return keymngr.Key
}

