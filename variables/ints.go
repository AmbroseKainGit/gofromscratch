package variables

import (
	"fmt"
)

func ShowInts() {
	// Variables are defined with the minimum value of the type for int is 0
	var integer int
	intde32 := int32(10)
	intde64 := int64(100)
	fmt.Println("integer equals = ", integer)
	fmt.Println("intde32 equals = ", intde32)
	fmt.Println("intde64 equals = ", intde64)
}
