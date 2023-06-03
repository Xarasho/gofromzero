package files

import (
	"fmt"
	"os"
	"bufio"

	"github.com/Xarasho/gofromzero/exercises"
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
	if !Append(fileName, text) {
		fmt.Println("Error while appending content")
	}
}

func Append(fileN string, text string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644 )
	if err != nil {
		fmt.Println("Error on append  " + err.Error())
		return false
	}

	_, err = arch.WriteString(text)
	if err != nil {
		fmt.Println("Error on WriteString  " + err.Error())
		return false
	}

	arch.Close()
	return true
}

func ReadFile() {
	archive, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error on ReadFile " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archive)
	for scanner.Scan() {
		register := scanner.Text()
		fmt.Println("> "+register)
	}
	archive.Close()
}
