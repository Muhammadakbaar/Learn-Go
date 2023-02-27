package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka = ", i)

	}
	var j = 0

	for j < 5 {
		fmt.Println("Angka2 ", j)
		j++
	}

	var k = 0
	for {
		fmt.Println("Angka3 ", k)

		k++
		if k == 5 {
			break
		}
	}

	for l := 1; l <= 10; l++ {
		if l%2 == 1 {
			continue
		}
		if l > 8 {
			break
		}
		fmt.Println("Angka4 ", l)
	}
}
