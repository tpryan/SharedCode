/*
	Copyright 2015, Google, Inc. 
 Licensed under the Apache License, Version 2.0 (the "License"); 
 you may not use this file except in compliance with the License. 
 You may obtain a copy of the License at 
  
    http://www.apache.org/licenses/LICENSE-2.0 
  
 Unless required by applicable law or agreed to in writing, software 
 distributed under the License is distributed on an "AS IS" BASIS, 
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. 
 See the License for the specific language governing permissions and 
 limitations under the License.
*/
package hello

import (
	"fmt"
	"log"
	"net/http"

	"appengine"
	"appengine/memcache"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/json", jsonHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	item, err := memcache.Get(c, "secret")
	if err != nil {
		handleError(w, err)
	}

	fmt.Fprint(w, string(item.Value))
}

type Location struct {
	Name  string  `json:"name"`
	City  string  `json:"city"`
	State string  `json:"state"`
	Lat   float32 `json:"latitude"`
	Lon   float32 `json:"longitude"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	var item Location
	key := "secret2"
	_, err := memcache.JSON.Get(c, key, &item)
	if err != nil {
		handleError(w, err)
	}

	fmt.Fprintf(w, "%+v\n", item)

}

func handleError(w http.ResponseWriter, err error) {
	log.Print("Yo")
	log.Printf("%v\n", err)
	http.Error(w, err.Error(), http.StatusInternalServerError)

}
