package pointers

import "fmt"

func Demo1(number *int) {
	*number = *number + 1
	fmt.Println(*number)

}
func Demo2(numbers []int) {
	numbers[0] = 100
	fmt.Println(numbers[0])

}
