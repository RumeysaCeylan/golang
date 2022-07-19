package functions

func Add(n1 int, n2 int) int {
	x := n1 + n2
	return x

}
func SayHi(n1 string) string {

	return "merhaba " + n1
}
func Calculator(num1 int, num2 int) (int, int, int, float64) {
	x := num1 + num2
	y := num1 - num2
	z := num1 * num2
	k := float64(num1 / num2)

	return x, y, z, k

}
func CalcVariadic(nums ...int) int {
	toplam := 0
	for i := 0; i < len(nums); i++ {
		toplam = toplam + nums[i]
	}
	return toplam
}
