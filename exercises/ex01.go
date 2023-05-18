package exercises

import "strconv"

func ConvertNum(text string) (int, string) {
	number, err := strconv.Atoi(text)
	if err != nil {
		return 0, "There's an error" + err.Error()
	}
	if number > 100 {
		return number, "It's more than 100"
	} else {
		return number, "It's less than 100"
	}
}