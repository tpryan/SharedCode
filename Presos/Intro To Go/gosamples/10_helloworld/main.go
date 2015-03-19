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

type Destination struct {
	Name     string
	Language string
	Country  string
	Distance int
}

func main() {
	defer un(trace("main"))
	des := []Destination{
		{"San Francisco", "English", "USA", 0},
		{"Mexico City", "Spanish", "Mexico", 1877},
		{"Bejing", "Chinese", "China", 5901},
		{"Montreal", "French", "Canada", 2535},
		{"Paramaribo", "Dutch", "Suriname", 4728},
		{"Windhoek", "German", "Namibia", 9819},
		{"Moscow", "Russian", "Russia", 5865},
		{"Tokyo", "Japanese", "Japan", 5136},
		{"Athens", "Greek", "Greece", 6772},
		{"Rome", "Italian", "Italy", 6239},
		{"Seoul", "Korean", "South Korea", 5607},
		{"São Paulo", "Portuguese", "Brazil", 6479},
	}

	var wg sync.WaitGroup

	for _, d := range des {
		wg.Add(1)
		go func(d Destination) {
			defer wg.Done()
			greet(d)
		}(d)
	}

	wg.Wait()
}

func greet(d Destination) {
	time.Sleep(time.Duration(d.Distance) * time.Millisecond)
	greet, _ := helloworld.Greet(d.Language)
	fmt.Printf("In %s (%d miles away), they say: %s\n", d.Name, d.Distance, greet)
}

func trace(s string) (string, time.Time) {
	return s, time.Now()
}

func un(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println(s, "ElapsedTime in seconds:", endTime.Sub(startTime))
}
