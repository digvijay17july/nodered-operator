package controller

import (
	"github.com/digvijay17july/nodered-operator/pkg/controller/nodered"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, nodered.Add)
}
