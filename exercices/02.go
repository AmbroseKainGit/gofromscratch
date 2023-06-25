package exercices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num int
var err error

func readInput(prompt string) (int, error) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	num, err = strconv.Atoi(input)
	return num, err
}

func MultiplesOfNumber() {
	for {
		num, err = readInput("Input a number: ")
		if err != nil {
			fmt.Println("Wrong value please enter a valid number")
			continue
		}
		break
	}
	for i := 0; i <= 10; i++ {
		fmt.Println(num, " x ", i, "=", (num * i))
	}

}
