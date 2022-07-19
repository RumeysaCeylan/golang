package variables

import "fmt"

func Demo1() {
	var str string = "selam dunya ben massaka" //tek tırnak olmaz
	fmt.Println(str)

	var kdv int = 15
	fmt.Println(100 + 10*kdv/100)
	var kdv2 float32 = 15.6
	fmt.Println(kdv2 * 10)

	kdv3 := 25.2
	fmt.Printf("%T\n", kdv3)

	var durum bool = true
	var durum2 bool
	var str1 string = "abc"
	var str2 string = "abc"
	durum = str1 == str2
	durum2 = str1 != str2

	fmt.Println(durum)

	if durum {
		fmt.Println(durum2)
		fmt.Println(!durum2)

	}

}
