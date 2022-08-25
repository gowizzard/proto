package build_test

import (
	"github.com/gowizzard/proto/build"
	"github.com/gowizzard/proto/dye"
	"testing"
	"time"
)

// TestLog is to test the build log function
// With table driven testing and logging of the returned values
func TestLog(t *testing.T) {

	tests := []struct {
		Kind       string
		Timestamp  time.Time
		Message    string
		Attributes map[string]any
		Color      bool
		Dye        dye.Config
	}{
		{
			Kind:       "info",
			Timestamp:  time.Now(),
			Message:    "This is an informational message.",
			Attributes: map[string]any{},
			Color:      true,
			Dye: dye.Config{
				Color: "cyan",
			},
		},
		{
			Kind:      "info",
			Timestamp: time.Now(),
			Message:   "This is an informational message without color.",
			Attributes: map[string]any{
				"cache": "The cache was successfully cleared",
			},
			Color: false,
			Dye:   dye.Config{},
		},
		{
			Kind:      "warning",
			Timestamp: time.Now(),
			Message:   "This is an warning message.",
			Attributes: map[string]any{
				"warning": "the memory space is full",
			},
			Color: true,
			Dye: dye.Config{
				Color: "yellow",
			},
		},
		{
			Kind:      "error",
			Timestamp: time.Now(),
			Message:   "This is an error message.",
			Attributes: map[string]any{
				"error": "wrong type",
			},
			Color: true,
			Dye: dye.Config{
				Color: "red",
			},
		},
	}

	for _, value := range tests {

		c := build.Config{
			Kind:       value.Kind,
			Timestamp:  value.Timestamp,
			Message:    value.Message,
			Attributes: value.Attributes,
			Dye:        value.Dye,
			Color:      value.Color,
		}
		log := c.Log()

		t.Logf(string(log))

	}

}
