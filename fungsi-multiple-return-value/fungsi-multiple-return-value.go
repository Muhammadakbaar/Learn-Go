package main

import "fmt"

func main() {
	firstname, lastnama := sayHai("", "")

	fmt.Println(firstname, lastnama)
	fmt.Println(sayHai("Sitamvan", " Sekale"))

}

func sayHai(firstname string, lastname string) (string, string) {
	if lastname == "" && lastname == "" {
		return "Muhammad", " Akbar"
	} else {
		return firstname, lastname
	}
}
