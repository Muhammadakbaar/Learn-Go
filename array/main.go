package main

import "fmt"

func main() {

	// array
	var nama [4]string

	nama[0] = "sitamvan"
	nama[1] = "muhammad"
	nama[2] = "akbar"
	nama[3] = "hakim"

	fmt.Println("=== Array ===")
	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])
	fmt.Println(nama[3])

	// inisialisasi nilai awal array
	var fruits = [4]string{"apple", "mangga", "durian", "manggis"}

	fmt.Println("=== Array dengan inisialisasi nilai awal ===")
	fmt.Println("Jumlah elemen \t\t", len(fruits))
	fmt.Println("Isi semua elemen \t", fruits)

	// array dengan gaya vertikal
	var fruits2 = [4]string{
		"apple",
		"mangga",
		"durian",
		"manggis",
	}
	fmt.Println("=== Array dengan gaya vertika ===")
	fmt.Println(fruits2)

	//  inisialisasi nilai awal array tanpa jumlah elemen
	var number2 = [...]int{2, 3, 2, 4, 3}

	fmt.Println("=== Array dengan inisialiasasi nilai awal tanpa jumlah elemen ===")
	fmt.Println("data array \t:", number2)
	fmt.Println("jumlah elemen \t:", len(number2))

	// array multidimensi
	var number3 = [3][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}, [3]int{4, 2, 5}}
	var number4 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("=== Array multidimensi ===")
	fmt.Println("number3 = ", number3, "jumlah elemen", len(number3))
	fmt.Println("number4 = ", number4)

	// array dengan perulangan for
	var fruits3 = [4]string{
		"apple",
		"manggo",
		"papaya",
		"batatas",
	}
	fmt.Println("=== Array dengan perulangan for ===")
	for i := 0; i < len(fruits3); i++ {
		fmt.Printf(" elemen %d : %s\n", i, fruits3[i])
	}

	// elemen array menggunakan keyword for - range
	fmt.Println("=== Array with for - range ===")
	var fruits4 = [4]string{"apple", "manggo", "sirsak", "banana"}
	for j, fruits5 := range fruits4 {
		fmt.Printf(" elemen %d : %s\n", j, fruits5)
	}

	// penggunaan variabel underscore ( _ )
	fmt.Println("=== penggunaan variabel ( _ ) ===")
	var fruits5 = [4]string{"apple", "manggo", "jeruk", "manggis"}

	for _, fruits6 := range fruits5 {
		fmt.Printf("nama buah %s \n", fruits6)
	}
	// penggunaan array dengan keyword make
	fmt.Println("=== Array menggunakan keyword make ===")

	var fruits7 = make([]string, 2)
	fruits7[0] = "apple"
	fruits7[1] = "manggo"

	fmt.Println(fruits7)
}
