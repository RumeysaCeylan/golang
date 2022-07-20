package errorhandling

import (
	"fmt"
	"os"
)

//TYPE ASSERTION
func Demo1() {
	f, err := os.Open("demo12.txt")
	//dosya varsa err nil olur yani error yoktur
	if err != nil {
		if pErr, k := err.(*os.PathError); k {
			fmt.Println("Dosya bulunamadı ", pErr.Path)
		} else {
			fmt.Println("Error")
			return
		}

	} else {
		fmt.Println(f.Name())
	}
}
