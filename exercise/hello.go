package main

func main() {
	//Variable declaration
	//	fmt.Println("Hello World")
	//
	//	var firstName = "enigma"
	//	fmt.Printf("Hello %s\n", firstName)
	//	firstName = "Camp"
	//	fmt.Printf("Hello %s\n", firstName)
	//
	//	lastName := "Camp"
	//	fmt.Printf("Hello %s %s\n", firstName, lastName)
	//
	//	var address string
	//	fmt.Printf("Alamat : %s\n", address)
	//
	//	var age int
	//	fmt.Printf("Alamat : %d\n", age)
	//
	//	var isActive bool
	//	fmt.Printf("Active : %v, tipe datanya : %T\n", isActive, isActive)
	//
	//	message := `
	//Syarat & Ketentuan
	//-------------------
	//  1. Minimal pembelian 0.1 gram
	//  2. Ada biaya transaksi sebesar Rp 2000,00
	//	`
	//	fmt.Println(message)
	//
	//	const (
	//		Male = iota
	//		Female
	//	)
	//	fmt.Println(Male, Female)
	//
	//	var(
	//		nama string
	//		usia int
	//	)
	//	fmt.Println(nama,usia)

	// User Input
	//var name string
	//fmt.Printf("dbHost : ")
	//fmt.Scanf("%s", &name)
	//
	//fmt.Printf("Hello %s\n", name)
	//
	//fmt.Printf("Home Address : ")
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//homeAddress := scanner.Text()
	//fmt.Println("Homebase:\n", homeAddress)
	//argsWithProg := os.Args
	//fmt.Println(argsWithProg)
	//dbHostArg := argsWithProg[1]
	//dbUSerArg := argsWithProg[2]
	//
	//fmt.Println(dbHostArg)
	//fmt.Println(dbUSerArg)

	//fullNameArg := flag.String("name", "", "User full name")
	//homeAddressArg := flag.String("address", "", "User home address")
	//isGoldMemberArg := flag.Bool("goldm", false, "If user is a gold member")
	//flag.Parse()
	//
	//fmt.Println("Full Name: ", *fullNameArg)
	//fmt.Println("Home Address: ", *homeAddressArg)
	//fmt.Println("Gold Member:", *isGoldMemberArg)
	//
	//fmt.Println("Gender:", os.Getenv("GENDER"))

	//Collection
	//var a [5]int
	//fmt.Println("int val", a)
	//
	//a[4] = 100
	//fmt.Println("int val", a)
	//
	//b := [...]int{1, 2, 3, 4, 5}
	//fmt.Printf("Array Default Value: %v, %T \n", b, b)
	//
	//var numbers = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	//fmt.Println("numbers: ", numbers)
	//
	//for _, intVal := range b {
	//	fmt.Printf("elemen with value %d\n", intVal)
	//}
	//for i := 0; i < len(b); i++ {
	//	fmt.Println(b[i])
	//}

	//for intVal := range b{
	//	fmt.Printf("value %d\n", intVal)
	//}

	//arrFood := [...]string{"Combro", "Gado2", "Mie ayam", "Tahu bulat", "Cilok", "Pisang goreng"}
	//fmt.Println("=== Menu Makanan ===")
	//fmt.Println(arrFood[3:])
	//food1 := arrFood[0:3]
	//fmt.Printf("%T %v \n", food1, food1)
	//fmt.Printf("Food 1 => address:%p, %v,len:%d,cap:%d\n", food1, food1, len(food1), cap(food1))
	//food1 = append(food1, "Roti bakar")
	//fmt.Printf("Food 1 => address:%p, %v,len:%d,cap:%d\n", food1, food1, len(food1), cap(food1))
	//
	//var sayuran []string
	//fmt.Printf("Sayuran => address:%p, %v,len:%d,cap:%d\n", sayuran, sayuran, len(sayuran), cap(sayuran))
	//sayuran = append(sayuran, "Kol")
	//fmt.Printf("Sayuran => address:%p, %v,len:%d,cap:%d\n", sayuran, sayuran, len(sayuran), cap(sayuran))
	//sayuran = append(sayuran, "bayam")
	//fmt.Printf("Sayuran => address:%p, %v,len:%d,cap:%d\n", sayuran, sayuran, len(sayuran), cap(sayuran))
	//
	//var buah = []string{"apple", "grape", "banana", "melon"}
	//fmt.Println(buah[0])
	//
	//var buah2an = make([]string, 0)
	//fmt.Printf("Buah2an => address:%p, %v,len:%d,cap:%d\n", buah2an, buah2an, len(buah2an), cap(buah2an))
	//buah2an = append(buah2an, "Jeruk")
	//fmt.Printf("Buah2an => address:%p, %v,len:%d,cap:%d\n", buah2an, buah2an, len(buah2an), cap(buah2an))
	//buah2an = append(buah2an, "Manggis")
	//fmt.Printf("Buah2an => address:%p, %v,len:%d,cap:%d\n", buah2an, buah2an, len(buah2an), cap(buah2an))
	//
	//m := make(map[string]string)
	//m["db_host"] = "localhost"
	//m["db_user"] = "root"
	//m["db_password"] = "P@ssw0rd"
	//fmt.Println("db map:", m)
	//
	//api := map[string]string{
	//	"base_url": "http://localhost",
	//	"port":     "8080",
	//}
	//fmt.Println("api map:", api)
	//v, ok := api["port"]
	//fmt.Println("The value:", v, "Present?", ok)

	//res := plus(1, 2)
	//fmt.Println(res)
	//
	//res, isEven := evenNumberChecker(2)
	//fmt.Println(res, isEven)
	//
	//sum(1, 2, 46, 32, 357, 323, 234)
	//
	//multiply := func(num1, num2 int) int {
	//	return num1 * num2
	//}
	//fmt.Printf("%T", multiply)
	//result := multiply(2, 3)
	//fmt.Println(result)

	//isEvenFilter := func(n int) bool {
	//	if n%2 == 0 {
	//		return true
	//	} else {
	//		return false
	//	}
	//}

	//isFizzFilter := func(n int) bool {
	//	if (n%3 == 0) && (n%5 != 0) {
	//		return true
	//	} else {
	//		return false
	//	}
	//}
	//isBuzzFilter := func(n int) bool {
	//	if (n%5 == 0) && (n%3 != 0) {
	//		return true
	//	} else {
	//		return false
	//	}
	//}
	//isFizzBuzzFilter := func(n int) bool {
	//	if (n%3 == 0) && (n%5 == 0) {
	//		return true
	//	} else {
	//		return false
	//	}
	//}
	//numbers := []int{1, 2, 3, 5, 6, 15}
	//resultFizzFilter := collectionFilter(numbers, isFizzFilter)
	//fmt.Println("Fizz:", resultFizzFilter)
	//resultBuzzFilter := collectionFilter(numbers, isBuzzFilter)
	//fmt.Println("Buzz:", resultBuzzFilter)
	//resultFizzBuzzFilter := collectionFilter(numbers, isFizzBuzzFilter)
	//fmt.Println("Fizz-Buzz:", resultFizzBuzzFilter)
	//
	//hasil := plusNamed(2, 3)
	//fmt.Println(hasil)
	//
	//fizzBuzzFilter := func(n int) string {
	//	if n%3 == 0 && n%5 == 0 {
	//		return "fizzbuzz"
	//	} else if n%3 == 0 {
	//		return "fizz"
	//	} else {
	//		return "buzz"
	//	}
	//}
	//resultFizz, resultBuzz, resultFizzBuzz := collectionFizzBuzzFilter(numbers, fizzBuzzFilter)
	//fmt.Println("Fizz:", resultFizz)
	//fmt.Println("Buzz:", resultBuzz)
	//fmt.Println("Fizz-Buzz:", resultFizzBuzz)
}

//func collectionFizzBuzzFilter(nums []int, filter func(int) string) (f []int, b []int, fb []int) {
//	for _, e := range nums {
//		res := filter(e)
//		switch res {
//		case "fizzbuzz":
//			fb = append(fb, e)
//			break
//		case "fizz":
//			f = append(f, e)
//			break
//		default:
//			b = append(b, e)
//		}
//	}
//	return
//}
//
//func collectionFilter(nums []int, filter func(int) bool) []int {
//	result := make([]int, 0)
//	for _, each := range nums {
//		if filtered := filter(each); filtered {
//			result = append(result, each)
//		}
//	}
//	return result
//}
//func plusNamed(a int, b int) (hasil int) {
//	hasil = a + b
//	return
//}
//func plus(a int, b int) int {
//	return a + b
//}
//
//func evenNumberChecker(num int) (int, bool) {
//	if num%2 == 0 {
//		return num, true
//	} else {
//		return -1, false
//	}
//}
//
//func sum(nums ...int) {
//	total := 0
//	for _, num := range nums {
//		total += num
//	}
//	fmt.Println(total)
//}
