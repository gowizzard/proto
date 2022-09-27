// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package build is there to build the log entry and
// populate it with all the information like kind,
// timestamp, message and attributes and return it.
package build

import (
	"github.com/gowizzard/proto/dye"
	"time"
)

// Config is to save essential data for log
// function and save data for other packages.
type Config struct {
	Kind       string
	Time       time.Time
	Message    any
	Attributes map[string]any
	Color      bool
	Dye        dye.Config
	Build      []byte
}
