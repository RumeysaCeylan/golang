package deferstatement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) Save() {
	fmt.Println("ürün kaydedildi : ", p.productName)
	defer Log()
}
func Log() {
	fmt.Println("Loglandı")
}
func Demo3() {
	p := product{productName: "laptop", unitPrice: 5000}
	defer p.Save()
	p = product{productName: "mouse", unitPrice: 45}

	fmt.Println("işlem başarılı")
}
