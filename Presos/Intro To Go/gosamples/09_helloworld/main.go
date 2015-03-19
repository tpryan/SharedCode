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
	"log"
	"sync"
	"time"
)

import (
	"github.com/tpryan/gosamples/helloworld"
)

// func main() {
// 	defer un(trace("main"))
// 	greetLocal()
// 	greetMexico()
// 	greetFrance()
// 	greetChina()
// }

// func main() {
// 	defer un(trace("main"))
// 	go greetLocal()
// 	go greetMexico()
// 	go greetChina()
// 	go greetFrance()
// 	time.Sleep(time.Second * 7)
// }

func main() {
	defer un(trace("main"))
	var wg sync.WaitGroup

	wg.Add(4)
	go func() {
		greetLocal()
		wg.Done()
	}()
	go func() {
		greetMexico()
		wg.Done()
	}()
	go func() {
		greetFrance()
		wg.Done()
	}()
	go func() {
		greetChina()
		wg.Done()
	}()
	wg.Wait()
}

func greetLocal() {
	defer un(trace("greetLocal"))
	greet, _ := helloworld.Greet("English")
	fmt.Printf("Here we say: %v\n", greet)
}

func greetMexico() {
	defer un(trace("greetMexico"))
	time.Sleep(time.Millisecond * 1877)
	greet, _ := helloworld.Greet("Spanish")
	fmt.Printf("In Mexico City (1877 miles away), they say: %v\n", greet)
}

func greetFrance() {
	defer un(trace("greetFrance"))
	time.Sleep(time.Millisecond * 5560)
	greet, _ := helloworld.Greet("French")
	fmt.Printf("In France (5560 miles away) they say: %v\n", greet)
}

func greetChina() {
	defer un(trace("greetChina"))
	time.Sleep(time.Millisecond * 5901)
	greet, _ := helloworld.Greet("Chinese")
	fmt.Printf("In China (5901 miles away) they say: %v\n", greet)
}

func trace(s string) (string, time.Time) {
	return s, time.Now()
}

func un(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println(s, "ElapsedTime in seconds:", endTime.Sub(startTime))
}
