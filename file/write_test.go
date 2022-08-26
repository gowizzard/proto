package file_test

import (
	"github.com/gowizzard/proto/file"
	"testing"
)

// TestWrite is to test the write function
// With table driven tests
func TestWrite(t *testing.T) {

	tests := []struct {
		Path string
		Log  []byte
	}{
		{
			Path: "/var/log/proto",
			Log:  []byte("This is a test log."),
		},
	}

	for _, value := range tests {

		c := file.Config{
			Path: value.Path,
		}

		err := c.Write(value.Log)
		if err != nil {
			t.Fatal(err)
		}

	}

}
