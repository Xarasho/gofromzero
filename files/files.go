package files

import (
	"fmt"
	"os"

	"github.com/Xarasho/gofromzero/exercises"
	//"bufio"
	//"ioutils"
)

var fileName string = "./files/txt/table.txt"

func StoreTable() {
	var text string = exercises.MultiplyInput()
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating file " + err.Error())
		return
	}
	fmt.Fprintln(file, text)
	file.Close()
}


func SumTable() {
	var text string = exercises.MultiplyInput()
	if Append(fileName, text) == false {
		fmt.Println("Error while appending content")
	}
}