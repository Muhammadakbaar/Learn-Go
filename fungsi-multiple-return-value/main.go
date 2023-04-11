package main

import "fmt"

func main() {

	fmt.Println(name("akbar", 20))

}

func name(firstname string, umur int) (string, int) {

	return firstname, umur

}
