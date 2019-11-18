// Package cross is the package keeping all the reusable cross cuting concern functionality
package cross

import (
	"fmt"
	"os"
)

// Error is the function to throw error to client
func Error(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(0)
}

// Print is the function to print output to client
func Print(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf(format, a...))
}
