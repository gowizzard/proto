package file

import (
	"time"
)

// Config is to save essential data for file package
type Config struct {
	Time    time.Time
	Monthly bool
	Path    string
	Name    string
}
