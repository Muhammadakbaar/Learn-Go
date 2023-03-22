package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumreference = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t:%.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %2f \n", circumreference)

	fmt.Println(iseng("menteng"))

}
func calculate(d float64) (float64, float64) {
	fmt.Println("=== fungsi dengan multiple return ===")
	var area = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d

	return area, circumference
}

func iseng(nama string) (namadepan string, namabelakang string) {

	namadepan = nama
	namabelakang = nama

	return
}
