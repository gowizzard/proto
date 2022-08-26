package file

import (
	"time"
)

// Config is to save essential data for file package
type Config struct {
	Time time.Time
	Path string
	Name string
}
