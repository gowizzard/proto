package file

import (
	"io/fs"
	"time"
)

// Config is to save essential data for file package
type Config struct {
	Filesystem fs.FS
	Time       time.Time
	Path       string
	Name       string
}
