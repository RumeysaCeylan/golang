package maps

import "fmt"

func Demo1() {
	//key-value
	dict := make(map[string]string)
	dict["table"] = "masa"
	dict["book"] = "kitap"
	dict["pencil"] = "kalem"

	fmt.Println(dict["book"])
	fmt.Println(len(dict))
	delete(dict, "book")
	fmt.Println(dict)

	value, isExist := dict["table"]
	fmt.Println(value)
	fmt.Println(isExist)

	dict2 := map[string]string{"glass": "bardak"}
	fmt.Println(dict2)

}

//for range

func ForRange() {
	cities := []string{"ANKARA", "İSTANBUL", "İZMİR"}

	for _, v := range cities {
		fmt.Println(v)
	}
	/*
		for i, v := range cities {
			fmt.Println(v)
		}
	*/
}
func Workshop() {
	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total := 0
	for _, v := range v {
		if v%2 == 0 {
			continue
		} else {
			total = total + v
		}
	}
	fmt.Println(total)
}
func Workshop2() {
	dict2 := map[string]string{"glass": "bardak", "pencil": "kalem", "table": "masa"}
	for k, v := range dict2 {
		fmt.Print(k + " ")
		fmt.Println(v)

	}
}
