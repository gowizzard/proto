package proto_test

import (
	"github.com/gowizzard/proto"
	"testing"
)

// TestLog is to test the logging function
// With table driven testing
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
				CommandLine:     true,
				FileInformation: true,
			},
			Kind:    "error",
			Message: "This is an error message.",
			Attributes: map[string]any{
				"error": "wrong type",
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
