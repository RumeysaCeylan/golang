package deferstatement

import "fmt"

func Demo(num int) string {
	defer DeferFunc()

	if num%2 == 0 {
		return "çift"
	} else if num%2 != 0 {
		//fmt.Println("tek")
		return "tek"
	}
	return "not sure"
}
func Test() {
	x := Demo(9)
	fmt.Println(x)
}
func DeferFunc() {
	fmt.Println("Bitti")
}
