package errorhandling

import "fmt"

type borderException struct {
	parameter int
	message   string
}

func (b borderException) Error() string {
	return fmt.Sprintf("%d %s", b.parameter, b.message)
}
func Guess2(guess int) (string, error) {
	if guess < 1 || guess > 100 {
		return "", &borderException{parameter: guess, message: "Sınırların dışında"}
	}
	return "doğru", nil
}
