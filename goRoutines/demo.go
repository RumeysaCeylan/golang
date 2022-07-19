package goroutines

import "fmt"

func EvenNumbers() {
	for i := 0; i <= 10; i += 2 {
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Print("\n")
}
func OddNumbers() {
	for i := 1; i <= 10; i += 2 {
		fmt.Print(i)
		fmt.Print(" ")
	}
}
