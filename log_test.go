package proto_test

import (
	"github.com/gowizzard/proto"
	"os"
	"testing"
)

// TestLog is to test the logging function
// With table driven testing and temporary directory for file testing
func TestLog(t *testing.T) {

	tests := []struct {
		Config     proto.Config
		Kind       string
		Message    string
		Attributes map[string]any
	}{
		{
			Config: proto.Config{
				CommandLine: true,
			},
			Kind:       "info",
			Message:    "This is an informational message.",
			Attributes: map[string]any{},
		},
		{
			Config: proto.Config{
				Convert: &map[string]string{
					"custom": "purple",
				},
				CommandLine: true,
			},
			Kind:    "custom",
			Message: "This is a custom message.",
		},
		{
			Config: proto.Config{
				CommandLine: true,
				Information: true,
			},
			Kind:       "info",
			Message:    "This is another informational message.",
			Attributes: map[string]any{},
		},
		{
			Config: proto.Config{
				CommandLine: true,
				Information: true,
			},
			Kind:    "error",
			Message: "This is an error message.",
			Attributes: map[string]any{
				"error": "wrong type",
			},
		},
		{
			Config: proto.Config{
				Information: true,
				Monthly:     true,
				CommandLine: true,
				File:        true,
				Path:        os.TempDir(),
			},
			Kind:    "warning",
			Message: "The buffer must be emptied.",
			Attributes: map[string]any{
				"storage": "95 percent occupied",
			},
		},
	}

	for _, value := range tests {

		c := value.Config
		err := c.Log(value.Kind, value.Message, value.Attributes)
		if err != nil {
			t.Fatal(err)
		}

	}

}
