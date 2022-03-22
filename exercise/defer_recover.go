package main

import (
	"errors"
	"fmt"
)

//func main() {
//	defer func() {
//		if err := recover(); err!=nil{
//			fmt.Println("Error recover")
//		}
//	}()
//	defer fmt.Println("aplikasi selesai")
//	fmt.Println("Selamat datang")
//	errorHappening()
//}
//
//func errorHappening() {
//	defer fmt.Println("Fungsi selesai")
//	fmt.Println("Proses 1")
//	panic("Error")
//	fmt.Println("Proses 2")
//}

func main() {
	total, err := simpleCalculation()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Total : ", total)

	//tulis("1")
	//tulis(1)
	//tulis(true)
	//tulis([]int{12, 132})
}

//func tulis(data interface{}) {
//	fmt.Println(data, "1", true, []float64{1.0})
//}

func simpleCalculation() (total float64, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = errors.New("error")
			total = -1
		}
	}()
	for i := 0; i <= 10; i++ {
		if i == 3 {
			panic("Error from DB / Network is disconnected")
		}
		total++
	}
	return
}
