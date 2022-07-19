package loops

import "fmt"

func Demo2() {
	x := 17
	var guess int
	i := 0
	for x != guess {
		fmt.Scanln(&guess)
		i++
		if x > guess {
			fmt.Println("Daha büyük")
		} else if x < guess {
			fmt.Println("daha küçük")
		} else {
			fmt.Println("doğru bildiniz")
			fmt.Printf("%v tahminde bildiniz  ", i)
			if i <= 3 && i >= 1 {
				fmt.Println("süper")
			} else if i <= 6 && i >= 4 {
				fmt.Println("Güzel")
			} else {
				fmt.Println("fena değil")
			}

		}
	}

}
