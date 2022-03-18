package main

import (
	"enigmacamp.com/goacc/model"
	"fmt"
)

type customerAddress struct {
	address    string
	city       string
	postalCode string
}

type customer struct {
	fullName string
	customerAddress
}

func main() {
	//customer := customer{
	//	fullName: "",
	//	customerAddress: customerAddress{
	//		address:    "",
	//		city:       "",
	//		postalCode: "",
	//	},
	//}

	product1 := model.NewProduct("111", "Sabun")

	fmt.Printf("%v \n", product1)
	product1.SetProductCode("XXX")
	fmt.Println(product1.ToString())

	product2 := model.Product{}
	product2.SetProductCode("222")
	product2.SetProductName("Sampo")
	fmt.Printf("%v \n", product2.ToString())

	//product3 := product{
	//	productCode: "333",
	//	productName: "Odol",
	//}
	//fmt.Printf("%v \n", product3)

}
