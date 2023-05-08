package main

import "fmt"

type Filter func(string) string

func main() {
	sayHello("Akbar", spamFilter)
	sayHello("Bangsat", spamFilter)

}
func sayHello(name string, filter Filter) {
	namedfilter := filter(name)
	fmt.Println("Hello", namedfilter)
}

func spamFilter(nama string) string {
	if nama == "Bangsat" {
		return "B***at"
	} else {
		return nama
	}
}
