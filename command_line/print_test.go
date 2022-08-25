package command_line_test

import (
	"github.com/gowizzard/proto/command_line"
	"testing"
)

// TestPrint is to test the printing function
// With table driven testing
func TestPrint(t *testing.T) {

	tests := []struct {
		Log []byte
	}{
		{
			Log: []byte("This is the first log."),
		},
		{
			Log: []byte("This is the second log."),
		},
		{
			Log: []byte("This is the third log."),
		},
	}

	for _, value := range tests {
		command_line.Print(value.Log)
	}

}
