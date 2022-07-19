package slices

import "fmt"

func Demo() {
	names := make([]string, 3)
	fmt.Println(names)
	names[0] = "ac"
	names[1] = "sdf"
	names[2] = "cfg"

	names = append(names, "rumeysa")
	fmt.Println(names)

	cities := []string{"ANKARA", "İSTANBUL"}
	fmt.Println(cities)
	copyCities := make([]string, len(cities))
	copy(copyCities, cities)
	fmt.Println(copyCities)
	cities = append(cities, "Adana")
	fmt.Println(cities)
	fmt.Println(copyCities)

	fmt.Println(cities[1:3])
	fmt.Println(cities[1:])
	fmt.Println(cities[:2])

}
