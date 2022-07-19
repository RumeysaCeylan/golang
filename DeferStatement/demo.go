package deferstatement

import "fmt"

func A() {
	fmt.Println("a")
}
func C() {
	fmt.Println("C")
}
func B() {
	fmt.Println("b")
	defer A() //önce b metodunu bitir sonra a yı çalıştır demektir ama belleğe a yı önce atar sonra c yi atar çıkarırken first come last out mantığında çalışır
	defer C()

	fmt.Println("b")

}
