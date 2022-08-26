package build

import (
	"github.com/gowizzard/proto/dye"
	"time"
)

// Config is to save essential data for log function
// And save data for other packages
type Config struct {
	Kind       string
	Time       time.Time
	Message    string
	Attributes map[string]any
	Color      bool
	Dye        dye.Config
	Build      []byte
}
