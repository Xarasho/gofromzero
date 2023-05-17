package variables

import (
	"fmt"
	"time"
) 

var Name string
var State bool
var Salary float32
var Date time.Time

func RestofVariables() {
	Name = "Pedro"
	State = true
	Salary = 1577.66
	Date = time.Now()	

	fmt.Println(Name)
	fmt.Println(State)
	fmt.Println(Salary)
	fmt.Println(Date)
}