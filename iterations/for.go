package iterations

import (
	"fmt"
)

func Iterate() {
	for i:=100; i>=0; i-=10 {
		if i == 60 {
			continue
		}
		fmt.Println(i)
	}
}