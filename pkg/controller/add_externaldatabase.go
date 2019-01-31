package controller

import (
	"github.com/UnlawfulMonad/edb-operator/pkg/controller/externaldatabase"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, externaldatabase.Add)
}
