package proto

import (
	"github.com/gowizzard/proto/build"
	"github.com/gowizzard/proto/file"
)

// Config is to save the client configuration
type Config struct {
	Information bool
	Convert     *map[string]string
	Monthly     bool
	CommandLine bool
	File        bool
	Path        string
	build       build.Config
	file        file.Config
}
