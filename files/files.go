package files

import (
	"fmt"
	"os"

	"github.com/Xarasho/gofromzero/exercises"
	//"bufio"
	//"ioutils"
)

func StoreTable() {
	var text string = exercises.MultiplyInput()
	file, err := os.Create("./files/txt/table.txt")
	if err != nil {
		fmt.Println("Error while creating file " + err.Error())
		return
	}
	fmt.Fprintln(file, text)
	file.Close()
}

func SumTable() {
	var text string = exercises.MultiplyInput()

}