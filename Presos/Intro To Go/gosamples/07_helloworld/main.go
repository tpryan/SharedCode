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
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/tpryan/gosamples/helloworld"
)

func main() {
	lang := flag.String("lang", "English", "a language in which to get greeting.")
	flag.Parse()

	greet, err := helloworld.Greet(*lang)

	if err != nil {
		if strings.Contains(err.Error(), helloworld.ErrorNotFound) {
			listLanguages()
		} else {
			log.Fatal(err)
		}
	}
	fmt.Printf("%s\n", greet)

}

func listLanguages() {
	list := helloworld.Langauges()
	fmt.Printf("Valid lanugages are: \n %v\n", list)
}
