// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package dye

import "fmt"

// start, reset are to generate the color code.
const (
	start = "\033["
	reset = "\033[0m"
)

// color is to map the color name to the code.
var (
	color = map[string]string{
		"red":    "31m",
		"green":  "32m",
		"yellow": "33m",
		"blue":   "34m",
		"purple": "35m",
		"cyan":   "36m",
	}
)

// Execute is to dye the text and return it as byte slice.
func (c Config) Execute(text string) []byte {

	c.Build = fmt.Append(c.Build, start)
	c.Build = fmt.Append(c.Build, color[c.Color])
	c.Build = fmt.Append(c.Build, text)
	c.Build = fmt.Append(c.Build, reset)

	return c.Build

}
