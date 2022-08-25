package command_line

import "fmt"

// Print is to print the log as string to the command line
func Print(log []byte) {
	fmt.Printf("%s\n", log)
}
