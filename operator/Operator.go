package main

import "fmt"

func main() {

	var angka1 int = 5
	var angka2 int = 5

	angka3 := angka1 + angka2
	angka4 := angka1 - angka2
	angka5 := angka1 * angka2
	angka6 := angka1 / angka2
	angka7 := angka1 % angka2

	fmt.Println("===Penjumlahan===")
	fmt.Println("5 + 5 = ", angka3)
	fmt.Println("===Pengurangan===")
	fmt.Println("5 - 5 = ", angka4)
	fmt.Println("===Perkalian===")
	fmt.Println("5 x 5 = ", angka5)
	fmt.Println("===Pembagian===")
	fmt.Println("5 / 5 = ", angka6)
	fmt.Println("===Modulus===")
	fmt.Println("5 % 5 = ", angka7)

}
