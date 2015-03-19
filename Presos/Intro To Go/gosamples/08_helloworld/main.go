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
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tpryan/gosamples/helloworld"
)

func main() {
	http.HandleFunc("/get", HandleGet)
	http.HandleFunc("/list", HandleList)
	log.Print(http.ListenAndServe(":8080", nil))
}

func HandleList(w http.ResponseWriter, r *http.Request) {
	list := helloworld.Langauges()

	json, err := json.Marshal(list)
	if err != nil {
		handleError(w, err)
		return
	}

	sendJSON(w, string(json))
}

func HandleGet(w http.ResponseWriter, r *http.Request) {

	lang := r.FormValue("language")

	greet, err := helloworld.Greet(lang)
	if err != nil {
		handleError(w, err)
		return
	}

	sendJSON(w, greet)
}

func handleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	log.Print(err)
}

func sendJSON(w http.ResponseWriter, content string) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, content)
}
