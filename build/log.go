package build

import (
	"fmt"
	"sort"
	"strings"
)

// Log is to build the log string
// We build the log string with a kind, timestamp, message and attributes
// To build the log for files, we remove the dye function for coloring the text
func (c Config) Log() []byte {

	c.Kind = strings.ToUpper(c.Kind)

	color := string(c.Dye.Execute(c.Kind))
	if c.Color {
		c.Kind = color
	}

	c.Kind = fmt.Sprintf("%s[%s]", c.Kind, c.Timestamp.Format("2006-01-02T15:04:05"))

	c.Build = fmt.Appendf(c.Build, "%-*s", len(c.Kind)+2, c.Kind)
	c.Build = fmt.Append(c.Build, c.Message)

	keys := make([]string, 0, len(c.Attributes))
	for index := range c.Attributes {
		keys = append(keys, index)
	}
	sort.Strings(keys)

	for _, value := range keys {

		key := strings.ToUpper(value)

		color = string(c.Dye.Execute(key))
		if c.Color {
			key = color
		}

		c.Build = fmt.Appendf(c.Build, "\t%s=%v", key, c.Attributes[value])

	}

	return c.Build

}
