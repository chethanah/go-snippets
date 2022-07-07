//Package main is the main routine
package main

import (
	"fmt"

	"github.com/chethanah/go-snippets/go-docs/somepackage"
	/*
		Alias - give a custom name to pkg in this scope
		mypackage "github.com/chethanah/go-snippets/go-packages/somepackage"
	*/)

//main is the main routine
func main() {
	somepackage.PrintPkg("Hello from package")
	fmt.Println("Hello from main")
}
