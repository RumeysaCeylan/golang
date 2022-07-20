package interfaces

import (
	"fmt"
)

func Validate(i interface{}) {
	number, ok := i.(int)

	fmt.Println(number, ok)
}
func Demo3() {
	var num interface{} = "abc"
	Validate(num)
	var num2 interface{} = 2
	Validate(num2)
}
