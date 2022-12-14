// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package file is used to output the event logs in
// a file, this file and also the folder structure
// are created automatically.
package file

import (
	"fmt"
	"os"
	"path/filepath"
)

// Write is to save the log to log file. We
// check the file and red the data, if exists.
// After that we write the byte slice to the existing data
func (c Config) Write(log []byte) error {

	log = fmt.Append(log, "\n")

	if c.Monthly {

		c.Name = c.Time.Format("2006-01") + ".log"
		c.Path = filepath.Join(c.Path, c.Name)

	} else {

		c.Path = filepath.Join(c.Path, c.Time.Format("2006-01"))
		err := os.MkdirAll(c.Path, 0750)
		if err != nil {
			return err
		}

		c.Name = c.Time.Format("2006-01-02") + ".log"
		c.Path = filepath.Join(c.Path, c.Name)

	}

	file, _ := os.ReadFile(c.Path)
	file = append(file, log...)

	err := os.WriteFile(c.Path, file, 0777)
	if err != nil {
		return err
	}

	return nil

}
