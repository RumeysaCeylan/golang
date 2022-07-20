package stringfunctions

import (
	"fmt"
	s "strings" //alias
)

func Demo1() {
	name := "Rümeysa"
	fmt.Println(s.Count(name, "y"))
	fmt.Println(s.Contains(name, "y"))
	fmt.Println(s.Index(name, "a"))
	fmt.Println(s.ToLower(name))
	fmt.Println(s.ToUpper(name))

}
