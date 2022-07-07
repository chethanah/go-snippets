package main

import (
	"fmt"

	"github.com/chethanah/go-snippets/go-packages/somepackage"
	/*
		Alias - give a custom name to pkg in this scope
		mypackage "github.com/chethanah/go-snippets/go-packages/somepackage"
	*/)

func main() {
	somepackage.PrintPkg("Hello from package")
	fmt.Println("Hello from main")
}
