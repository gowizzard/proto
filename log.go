// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package proto

import (
	"fmt"
	"github.com/gowizzard/proto/command_line"
	"path/filepath"
	"reflect"
	"runtime"
	"time"
)

// color are to save the color name.
// convert are to save the kind name to a color name.
var (
	color   string
	convert = map[string]string{}
)

// Log ist to print the message to command line
// or file. You can use the config struct to
// filter data and change the functionality.
func (c Config) Log(kind string, message any, attributes map[string]any) error {

	c.build.Kind = kind
	c.build.Time = time.Now()
	c.build.Message = message

	if c.Information {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			attributes["file"] = fmt.Sprintf("%s:%d", filepath.Base(file), line)
		}
	}
	c.build.Attributes = attributes

	switch c.Convert {
	case nil:
		convert = map[string]string{
			"default": "cyan",
			"error":   "red",
			"warning": "yellow",
		}
	default:
		convert = *c.Convert
	}

	color = convert[kind]
	if reflect.ValueOf(color).Len() == 0 {
		color = convert["default"]
	}
	c.build.Dye.Color = color

	if c.CommandLine {

		c.build.Color = true

		log := c.build.Log()
		command_line.Print(log)

	}

	if c.File {

		c.build.Color = false

		c.file.Time = c.build.Time
		c.file.Monthly = c.Monthly
		c.file.Path = c.Path

		log := c.build.Log()
		err := c.file.Write(log)
		if err != nil {
			return err
		}

	}

	return nil

}
