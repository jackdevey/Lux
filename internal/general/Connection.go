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

import (
	"bytes"
	"io"
	"net/http"
)

// Connection contains information about
// the connection to the Govee API such
// as the user's API Key & the server URL.
type Connection struct {
	http http.Client
	Key  string
	Url  string
}

// Get makes a GET request to the API server
// with the path provided as a string.
// Usually a JSON object is returned in
// byte array form to help with parsing.
func (c Connection) Get(path string) []byte {
	request, _ := http.NewRequest("GET", c.Url + path, nil)
	request.Header.Set("Govee-API-Key", c.Key)
	response, err := c.http.Do(request)

	if err != nil {
		println("Error making request")
		return []byte("")
	}else if response != nil {
		body, _ := io.ReadAll(response.Body)
		return body
	}else {
		println("Something bad happened :(")
		return []byte("")
	}
}

func (c Connection) Put(path string, body []byte) []byte {
	request, _ := http.NewRequest("PUT", c.Url + path, bytes.NewBuffer(body))
	request.Header.Set("Govee-API-Key", c.Key)
	response, err := c.http.Do(request)

	if err != nil {
		println("Error making request")
		return []byte("")
	}else if response != nil {
		body, _ := io.ReadAll(response.Body)
		return body
	}else {
		println("Something bad happened :(")
		return []byte("")
	}
}
