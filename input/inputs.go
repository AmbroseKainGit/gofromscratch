package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// var num1 int
// var num2 int
// var message string
// var err error

// func InputNumber() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Println("Input number 1: ")
// 	if scanner.Scan() {
// 		num1, err = strconv.Atoi(scanner.Text())
// 		if err != nil {
// 			panic("Wrong input: " + err.Error())
// 		}
// 	}
// 	fmt.Println("Input number 2: ")
// 	if scanner.Scan() {
// 		num2, err = strconv.Atoi(scanner.Text())
// 		if err != nil {
// 			panic("Wrong input: " + err.Error())
// 		}
// 	}
// 	fmt.Println("Input message 2: ")
// 	if scanner.Scan() {
// 		message = scanner.Text()
// 	}

// 	fmt.Println(message, (num1 * num2))
// }

func readInput(prompt string) (int, error) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	num, err := strconv.Atoi(input)
	return num, err
}

func InputNumber() {
	num1, err := readInput("Input number 1: ")
	if err != nil {
		fmt.Println("Wrong input:", err.Error())
		return
	}

	num2, err := readInput("Input number 2: ")
	if err != nil {
		fmt.Println("Wrong input:", err.Error())
		return
	}

	fmt.Print("Input message: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message := scanner.Text()

	result := num1 * num2
	fmt.Println(message, result)
}
