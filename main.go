package main

// "github.com/Xarasho/gofromzero/keyboard"
// "github.com/Xarasho/gofromzero/iterations"
// "fmt"
// "github.com/Xarasho/gofromzero/files"

import (
	"fmt"

	"github.com/Xarasho/gofromzero/exercises"
)

func main() {
	//variables.ShowIntegers()
	//variables.RestofVariables()
	// state, text := variables.ConverToText(1588)
	// fmt.Println(state)
	// fmt.Println(text)

	// if 	os := runtime.GOOS; os=="Linux."  || os=="OS X." {
	// 	fmt.Println("Esto no es Windows, es ", os)
	// } else{
	// 	fmt.Println("Esto es Windows")
	// } 

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }
		
	// number, text := exercises.ConvertNum("69")
	// fmt.Println(number)
	// fmt.Println(text)

	// keyboard.InputNumbers()

	// iterations.Iterate()

	fmt.Println(exercises.MultiplyInput())

	// files.WriteTable()

	// files.SumTable()
}