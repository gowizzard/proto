// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package command_line is used to output the information
// to the command line interface so that it is visible there.
package command_line

import "fmt"

// Print is to print the log as string to the command line.
func Print(log []byte) {
	fmt.Printf("%s\n", log)
}
