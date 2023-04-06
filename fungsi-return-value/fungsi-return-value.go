package main

import "fmt"

func sayHay(name string) string {
	if name == "" {
		return "Hello bro"
	} else {
		return "Hello" + name
	}

}
func main() {

	result := sayHay("")
	fmt.Println(result)
	fmt.Println(sayHay(" Akbar"))
}
