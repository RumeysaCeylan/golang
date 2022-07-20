package stringfunctions

import (
	"fmt"
	s "strings" //alias
)

func Demo2() {
	name := "Rümeysa"
	fmt.Println(s.HasPrefix(name, "Rüm"))
	fmt.Println(s.HasSuffix(name, "sa"))
	fmt.Println(s.Index(name, "ey"))
	alphbt := []string{"r", "m", "y", "s"}
	snc := s.Join(alphbt, "*")
	fmt.Println(snc)

	fmt.Println(s.Replace(snc, "*", "+", -1)) //-1 tüm hepsini değiştir demek ama sayı verirsek kaç tane görürse o kadar değişir
	fmt.Println(s.Split(snc, "*"))
	fmt.Print(s.Repeat(snc, 5))
}
