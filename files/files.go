package files

import (
	"os"

	"github.com/Xarasho/gofromzero/exercises"
	//"bufio"
	//"ioutils"
)

func StoreTable() {
	var text string = exercises.MultiplyInput()
	file, err := os.Create("./files/txt/table.txt")
}

func SumTable() {
	var text string = exercises.MultiplyInput()

}