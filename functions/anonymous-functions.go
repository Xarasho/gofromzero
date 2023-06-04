package functions

import "fmt"

func Calc() {

	sum := func(number1 int, number2 int) int {
		return number1 + number2
	}

	fmt.Println(sum(10, 25))
}