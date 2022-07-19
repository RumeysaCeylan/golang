package errorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, error := os.Open("demo1.txt")
	//dosya varsa err nil olur yani error yoktur
	if error != nil {
		fmt.Println("Dosya bulunamadı")
	} else {
		fmt.Println(f.Name())
	}
}
