// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package dye_test

import (
	"github.com/gowizzard/proto/dye"
	"testing"
)

// TestExecute is to test the dye execute function. We
// test the execute function via table driven tests.
func TestExecute(t *testing.T) {

	tests := []struct {
		Color string
		Text  string
	}{
		{
			Color: "red",
			Text:  "This is a red text test.",
		},
		{
			Color: "green",
			Text:  "This is a green text test.",
		},
		{
			Color: "yellow",
			Text:  "This is a yellow text test.",
		},
		{
			Color: "blue",
			Text:  "This is a blue text test.",
		},
		{
			Color: "purple",
			Text:  "This is a purple text test.",
		},
		{
			Color: "cyan",
			Text:  "This is a cyan text test.",
		},
	}

	for _, value := range tests {

		c := dye.Config{
			Color: value.Color,
		}
		text := c.Execute(value.Text)

		t.Log(string(text))

	}

}
