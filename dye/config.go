// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package dye is used to color text elements
// in one of the stored colors and return
// this modified string as a slice of bytes.
package dye

// Config is to save the essential data like the
// color. And the build value to return the data.
type Config struct {
	Color string
	Build []byte
}
