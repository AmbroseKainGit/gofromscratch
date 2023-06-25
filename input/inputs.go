package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var message string
var err error

func InputNumber() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input number 1: ")
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Wrong input: " + err.Error())
		}
	}
	fmt.Println("Input number 2: ")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Wrong input: " + err.Error())
		}
	}
	fmt.Println("Input message 2: ")
	if scanner.Scan() {
		message = scanner.Text()
	}

	fmt.Println(message, (num1 * num2))
}
