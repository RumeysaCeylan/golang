package errorhandling

import (
	"errors"
	"fmt"
)

func Guess(guess int) (string, error) {
	if guess < 1 || guess > 100 {
		return "", errors.New("1-100 arasında giriniz")
	}
	return "Doğru", nil
}
func Demo2() {
	fmt.Println(Guess(50))
	fmt.Println(Guess(101))
	fmt.Println(Guess(-10))

}
