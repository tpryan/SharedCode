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
	"fmt"
	"math/rand"
	"time"
)

func main() {
	greetings := []string{"你好世界",
		"Hello wereld",
		"Hello world",
		"Bonjour monde",
		"Hallo Welt",
		"γειά σου κόσμος",
		"Ciao mondo",
		"こんにちは世界",
		"여보세요 세계",
		"Olá mundo",
		"Здравствулте мир",
		"Hola mundo"}

	rand.Seed(time.Now().Unix())
	l := len(greetings)
	i := rand.Intn(l - 1)

	fmt.Printf("%s\n", greetings[i])

}
