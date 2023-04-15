package main

import "fmt"

func fullName() (firstName, middleName, lastName string, umur int) {
	firstName = "Muhammad "
	middleName = "Akbar "
	lastName = "Hakim"
	umur = 25

	return
}
func main() {
	a, b, c, d := fullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	e, f := campuran("muhammad", 25)
	fmt.Println(e)
	fmt.Println(f)

}
func campuran(nama string, umur int) (a string, b int) {
	a = "akbar"
	b = 30
	return a, b
}
