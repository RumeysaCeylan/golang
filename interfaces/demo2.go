package interfaces

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}
type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 12 / 100
}
func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 12 / 100
}
func MontlyPaymentCalculate(credits []CreditCalculator) float64 {
	paymetTotal := 0.0
	for i := 0; i < len(credits); i++ {
		paymetTotal = paymetTotal + credits[i].Calculate()
	}
	return paymetTotal
}
func Demo2() {
	credit1 := Mortgage{rate: 10, address: "ABD", creditPaymentTotal: 100000}
	credit2 := Mortgage{rate: 12, address: "UK", creditPaymentTotal: 50000}
	credit3 := Car{rate: 15, creditPaymentTotal: 60000, carInfo: "polo"}

	credits := []CreditCalculator{credit1, credit2, credit3}
	total := MontlyPaymentCalculate(credits)
	fmt.Println(total)

}
