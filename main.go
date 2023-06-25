// Everything start with package (go convention)
package main

import (
	"github.com/AmbroseKainGit/gofromscratch/variables"
)

// Function that are capitalized are by default global to the package and all the package that import that package

func main() {
	// fmt.Println("Hello world")
	// variables.ShowInts()
	variables.AllVariables()
}
