package overlay

import "github.com/araddon/qlbridge/expr/builtins"

func init()  {
	// load all of our built-in functions
	builtins.LoadAllBuiltins()
}