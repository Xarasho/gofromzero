package keyboard

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

var number1 int
var number2 int
var leyend string
var err error

func InputNumbers() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Input number 1: ")
	if scanner.Scan() {
		number1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Input data is incorrect" + err.Error())
		}
	}

	fmt.Println("Input number 2: ")
	if scanner.Scan() {
		number2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Input data is incorrect" + err.Error())
		}
	}

	fmt.Println("Input leyend: ")
	if scanner.Scan() {
		leyend = scanner.Text()
	}

	fmt.Println(leyend, number1*number2)
}