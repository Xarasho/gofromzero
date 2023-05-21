package exercises

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

var number int
var err error

func InputKeyboard() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input a number: ")
	
	if scanner.Scan() {
		number, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Please input a number" + err.Error())
		}

		for i:=1; i <= 10; i++ {
			fmt.Println(number*i)
		}
	}
}