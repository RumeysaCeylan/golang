package conditionals

import "fmt"

func Demo2() {
	var budget float64 = 1000
	var takenMoney float64 = 100

	var durum bool = budget > takenMoney

	if durum {
		fmt.Println("durum doğru")
	}

	if budget < takenMoney {
		fmt.Println("para yetersiz")
	} else if budget == takenMoney {
		fmt.Println("hesabınızdaki tüm parayı çektiniz")
	} else {
		budget = budget - takenMoney
		fmt.Println("hesapta kalan para " + fmt.Sprintf("%v", budget))
		fmt.Printf("hesaptaki para :%v ", budget)
	}

}
