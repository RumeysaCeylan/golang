package workshop1

import "fmt"

func Demo() {
	var num1, num2, num3 int = 40, 20, 30

	if num1 > num2 {
		if num1 > num3 {
			fmt.Println(num1)
		} else {
			fmt.Println(num3)
		}
	} else if num2 > num1 {
		if num2 > num3 {
			fmt.Println(num2)
		} else {
			fmt.Println(num3)
		}
	}
}
