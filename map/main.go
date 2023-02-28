package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["february"] = 40

	fmt.Println("januari ", chicken["januari"])
	fmt.Println("mei", chicken["mei"])
}
