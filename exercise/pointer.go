package main

import "fmt"

func main() {
	var number = 4
	fmt.Printf("Init: %v, tipe : %T, value: %p\n", number, number, &number)
	change(&number, 10)
	fmt.Printf("result Change: %v,tipe : %T, value: %p\n", number, number, &number)

	var myNamePtr = new(string)
	*myNamePtr = "enigma"
	salutation(myNamePtr)
	fmt.Printf("MyName pointer: %s, tipe : %T, address: %p\n", *myNamePtr, myNamePtr, myNamePtr)

}

func salutation(customer *string) {
	fmt.Printf("Customer Salutation: %v, tipe : %T, address: %p\n", *customer, customer, customer)
	*customer = fmt.Sprintf("Mr/Mrs/Ms. %s", *customer)
}
func change(original *int, newValue int) {
	fmt.Printf("Change: %p, tipe : %T, value: %v\n", original, original, *original)
	*original = newValue
}

func changeVal(original int, newValue int) {
	fmt.Printf("ChangeVal: %v, tipe : %T, value: %v\n", &original, original, original)
	original = newValue
}
