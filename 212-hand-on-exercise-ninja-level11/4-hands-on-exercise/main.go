// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := Sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

// function test go documentation
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your code here
		em := fmt.Errorf("norgate math redux: square root of negative number")

		// se := sqrtError{
		// 	lat:  "50.2289 N",
		// 	long: "99.4656 W",
		// 	err:  em,
		// }

		// return 0, se

		return 0, sqrtError{"50.2289 N", "99.4656 W", em}
	}
	return 42, nil
}
