package arrays

import "fmt"

func Demo() {
	var numbers [5]int
	numbers[2] = 50
	fmt.Println(numbers)
	fmt.Println(numbers[2])
	var cities [5]string
	cities[0] = "Ankara"
	cities[1] = "İstanbul"
	cities[2] = "izmir"
	cities[3] = "adana"
	cities[4] = "çorum"
	for i := 0; i < 5; i++ {
		//fmt.Println(cities[i])

	}
	nums := [5]int{20, 30, 2, 10, 5}
	for i := 0; i < len(nums); i++ {
		//fmt.Println(nums[i])

	}

	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[0][1] = 3
	matrix[0][2] = 5
	matrix[1][0] = 2
	matrix[1][1] = 4
	matrix[1][2] = 6
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}
