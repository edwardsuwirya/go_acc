package delivery

import (
	"fmt"
	"strings"
)

func MainMenu() {
	fmt.Println(strings.Repeat("*", 50))
	fmt.Println("Aplikasi")
	fmt.Println(strings.Repeat("*", 50))
	fmt.Println("1. Buat Produk Baru")
	fmt.Println("2. Daftar Produk")
	fmt.Println("3. Cari Produk")
	fmt.Println("4. Keluar")
	fmt.Println("Pilih menu (1-4): ")

}
