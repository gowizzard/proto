// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package file

import (
	"time"
)

// Config is to save essential data for file package.
type Config struct {
	Time    time.Time
	Monthly bool
	Path    string
	Name    string
}
