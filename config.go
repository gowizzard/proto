// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package proto is used to simplify and standardize
// logging. Several configurations can be selected for
// this. With this package you can perform logging
// in the cli as well as file-based logging.
package proto

import (
	"github.com/gowizzard/proto/build"
	"github.com/gowizzard/proto/file"
)

// Config is to save the client configuration.
type Config struct {
	Information bool
	Convert     *map[string]string
	Monthly     bool
	CommandLine bool
	File        bool
	Path        string
	build       build.Config
	file        file.Config
}
