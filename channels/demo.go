package channels

func EvenNumbers(evenNumCn chan int) {
	total := 0
	for i := 0; i <= 10; i += 2 {
		total += i
	}
	evenNumCn <- total
}
func OddNumbers(oddNumCn chan int) {
	total := 0
	for i := 1; i <= 10; i += 2 {
		total += i
	}
	oddNumCn <- total
}
