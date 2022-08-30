package file_test

import (
	"github.com/gowizzard/proto/file"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// TestWrite is to test the write function
// With table driven tests and temporary directory for file testing
func TestWrite(t *testing.T) {

	tests := []struct {
		Time    time.Time
		Monthly bool
		Path    string
		Log     []byte
	}{
		{
			Time: time.Now(),
			Path: os.TempDir(),
			Log:  []byte("This is a test log."),
		},
		{
			Time: time.Now().AddDate(-10, 0, -1),
			Path: os.TempDir(),
			Log:  []byte("This the second test log."),
		},
		{
			Time:    time.Now(),
			Monthly: true,
			Path:    os.TempDir(),
			Log:     []byte("This the third test log."),
		},
	}

	for _, value := range tests {

		c := file.Config{
			Time:    value.Time,
			Monthly: value.Monthly,
			Path:    value.Path,
		}

		err := c.Write(value.Log)
		if err != nil {
			t.Fatal(err)
		}

		var path, name string
		if value.Monthly {
			path = filepath.Join(c.Path)
			name = value.Time.Format("2006-01") + ".log"
		} else {
			path = filepath.Join(c.Path, value.Time.Format("2006-01"))
			name = value.Time.Format("2006-01-02") + ".log"
		}
		path = filepath.Join(path, name)

		read, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(read))

	}

}
