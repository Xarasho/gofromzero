package main

import (
	"fmt"

	"github.com/Xarasho/gofromzero/variables"
)

func main() {
	//variables.ShowIntegers()
	//variables.RestofVariables()
	state, text := variables.ConverToText(1588)
	fmt.Println(state)
	fmt.Println(text)
}