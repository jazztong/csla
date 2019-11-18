package generator

import (
	"os"

	"github.com/jazztong/csla/cross"
)

// Generator is the interface define the capability of template generator
type Generator interface {
	// Generate generate the template into user environment
	Generate(req Request)
}

// Validate is the function to validate if can generate template
func Validate(req Request) {
	if _, err := os.Stat(req.Name); !os.IsNotExist(err) {
		cross.Error("%s folder should not be empty", req.Name)
	}
}
