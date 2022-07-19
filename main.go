package main

import (
	errorhandling "golesson/errorHandling"
)

//"fmt"
//"golesson/arrays"
//"golesson/slices"
//"fmt"
//"golesson/functions"
//"golesson/maps"

//"golesson/conditionals"
//"golesson/loops"
//"golesson/variables" //go mod init golesson
//"golesson/workshop1"

func main() {

	//variables.Demo1()
	//fmt.Println(":)")
	//conditionals.Demo2()
	//workshop1.Demo()
	//loops.Demo2()
	//arrays.Demo()
	//slices.Demo()
	// x := functions.Add(2, 3)
	// y := functions.SayHi("rümeysa")
	// fmt.Println(x)
	// fmt.Println(y)
	// a, b, _, d := functions.Calculator(8, 5) //_değeri kullanmak istenmediğinde _ koyulur
	// fmt.Println(a)
	// fmt.Println(b)
	// //fmt.Println(c)
	// fmt.Println(d)
	// var s = functions.CalcVariadic(1, 2, 3, 4)
	// fmt.Println(s)
	// fmt.Println(functions.CalcVariadic(1, 2))
	// fmt.Println(functions.CalcVariadic())

	// nums := []int{4, 6, 7, 0, 11}
	// fmt.Println(functions.CalcVariadic(nums...))

	// maps.Demo1()
	// maps.ForRange()
	// maps.Workshop2()
	// num := 20
	// pointers.Demo1(&num)
	// fmt.Println(num)

	// numbers := []int{1, 2, 3}
	// pointers.Demo2(numbers)
	// fmt.Println(numbers[0])
	// structs.Demo2()
	// go goroutines.EvenNumbers()
	//go goroutines.OddNumbers()
	//time.Sleep(5 * time.Second)
	// evenNumCn := make(chan int)
	// oddNumCn := make(chan int)
	// go channels.EvenNumbers(evenNumCn)
	// go channels.OddNumbers(oddNumCn)
	// evenNumTotal, oddNumTotal := <-evenNumCn, <-oddNumCn
	// fmt.Println(evenNumTotal * osddNumTotal)

	// interfaces.Demo2()
	// deferstatement.B()
	// deferstatement.Test()
	// deferstatement.Demo3()
	errorhandling.Demo1()

}
