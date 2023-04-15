package main

import "fmt"

func sayHello(name string, filter func(string) string) {
	namedfilter := filter(name)
	fmt.Println("Hello", namedfilter)
}

func spamFilter(name string) string {
	if name == "Bangsat" {
		return "B***at"
	} else {
		return name
	}
}
func main() {
	sayHello("Akbar", spamFilter)
	sayHello("Bangsat", spamFilter)

}
