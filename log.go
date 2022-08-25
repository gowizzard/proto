package proto

import (
	"github.com/gowizzard/proto/command_line"
	"reflect"
	"time"
)

// color, convert are to save the color name
// And convert the kind name to a color name
var (
	color   string
	convert = map[string]string{
		"default": "cyan",
		"error":   "red",
		"warning": "yellow",
	}
)

// Log ist to print the message to command line or file
// You can use the config struct to filter data and change the functionality
func (c Config) Log(kind, message string, attributes map[string]any) error {

	c.build.Kind = kind
	c.build.Timestamp = time.Now()
	c.build.Message = message
	c.build.Attributes = attributes

	color = convert[kind]
	if reflect.ValueOf(color).Len() == 0 {
		color = convert["default"]
	}
	c.build.Dye.Color = color

	log := c.build.Log()

	if c.CommandLine {
		command_line.Print(log)
	}

	if c.File {
		c.file.Path = c.Path
		err := c.file.Write(log)
		if err != nil {
			return err
		}
	}

	return nil

}
