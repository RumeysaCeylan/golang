package loops

import "fmt"

func Demo() {
	x := 17
	var guess int
	i := 0
	for ; i < 3; i++ {
		fmt.Scanln(&guess)
		if x == guess {
			fmt.Println("doğru")
			break
		} else if x < guess {
			fmt.Println("Daha küçük")
		} else if x > guess {
			fmt.Println("Daha büyük")
		} else {
			fmt.Println("error")
		}

	}
	if i == 3 {
		fmt.Println("hakkınız doldu")

	}
	/*
		for x!=guess{
			//codes
		}

	*/

}
