package main

import "fmt"

func main() {
	sayHello("Muhammad", "Akbar")
}
func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)

}