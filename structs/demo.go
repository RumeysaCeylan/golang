package structs

import "fmt"

/*func Demo() {
	fmt.Println(product{"Bilgisayar", 500, "XYZ", 20})
	fmt.Println(product{name: "Bilgisayar", price: 500})

}

type product struct {
	name         string
	price        float64
	brand        string
	discountRate int
}*/
type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) Save() {
	fmt.Println("Eklendi ", c.firstName)
}
func Demo2() {
	c := customer{firstName: "Rümeysa", lastName: "Ceylan", age: 23}
	c.Save()
}
