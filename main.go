// Everything start with package (go convention)
package main

import (
	"fmt"
	"runtime"
)

// Function that are capitalized are by default global to the package and all the package that import that package

func main() {
	// fmt.Println("Hello world")
	// variables.ShowInts()
	// variables.AllVariables()
	// estado, text := variables.AnythingToText(1500)
	// fmt.Println(estado)
	// fmt.Println(text)
	// os := runtime.GOOS
	// First declare variable then do the validation
	if os := runtime.GOOS; os != "windows" {
		fmt.Println("This is not windosws")
	} else {
		fmt.Println("This is  windosws")
	}
	// if os := runtime.GOOS; os != "windows" {
	// 	fmt.Println("This is not windosws is : ", os)
	// 	return
	// }
	// fmt.Println("This is  windosws")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("linux")
	case "darwind":
		fmt.Println("darwind")
	default:
		fmt.Printf("%s \n", os)
	}

}
