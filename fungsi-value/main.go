package main

import "fmt"

func getName(name string) string {
	return "nama gue" + name
}
func main() {

	name := getName(" Muhammad")

	fmt.Println(getName(" Akbar"))
	fmt.Println(name)

}
