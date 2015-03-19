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
package helloworld

import "errors"

const (
	ErrorNotFound = "The language you selected is not valid."
)

var greetings = map[string]string{
	"Chinese":    "你好世界",
	"Dutch":      "Hello wereld",
	"English":    "Hello world",
	"French":     "Bonjour monde",
	"German":     "Hallo Welt",
	"Greek":      "γειά σου κόσμος",
	"Italian":    "Ciao mondo",
	"Japanese":   "こんにちは世界",
	"Korean":     "여보세요 세계",
	"Portuguese": "Olá mundo",
	"Russian":    "Здравствулте мир",
	"Spanish":    "Hola mundo",
}

func Greet(language string) (string, error) {

	if _, ok := greetings[language]; ok {
		return greetings[language], nil
	} else {
		return "", errors.New(ErrorNotFound)
	}

}

func Langauges() []string {
	var response []string
	for l := range greetings {
		response = append(response, l)
	}
	return response
}
