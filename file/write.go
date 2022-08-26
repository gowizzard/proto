package file

import (
	"fmt"
	"os"
	"path/filepath"
)

// Write is to save the log to log file
// We check the file and red the data, if exists
// After that we write the byte slice to the existing data
func (c Config) Write(log []byte) error {

	log = fmt.Append(log, "\n")

	c.Path = filepath.Join(c.Path, c.Time.Format("2006-01"))
	err := os.MkdirAll(c.Path, 0750)
	if err != nil {
		return err
	}

	c.Name = c.Time.Format("2006-01-02") + ".log"
	c.Path = filepath.Join(c.Path, c.Name)

	file, _ := os.ReadFile(c.Path)
	file = append(file, log...)

	err = os.WriteFile(c.Path, file, 0777)
	if err != nil {
		return err
	}

	return nil

}
