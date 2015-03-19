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

import (
	"errors"
	"testing"
)

func TestGreet(t *testing.T) {

	cases := []struct {
		input string
		s     string
		err   error
	}{
		{"", "", errors.New(ErrorNotFound)},
		{"English", "Hello world", nil},
	}

	for _, c := range cases {
		got_s, got_err := Greet(c.input)

		if got_s != c.s {
			t.Errorf("Validate(%q) == %q, want %q", c.input, got_s, c.s)
		}
		if got_err != nil && got_err.Error() != c.err.Error() {
			t.Errorf("Validate(%q) == %q, want %q", c.input, got_err, c.err)
		}

	}

}
