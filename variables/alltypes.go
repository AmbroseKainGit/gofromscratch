package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var State bool
var Price float32
var Date time.Time

func AllVariables() {
	Name = "Pedro"
	State = true
	Price = 150
	Date = time.Now()
	fmt.Println(Name)
	fmt.Println(State)
	fmt.Println(Price)
	fmt.Println(Date)
}

func AnythingToText(number int) (bool, string) {
	return true, strconv.Itoa(number)
}
