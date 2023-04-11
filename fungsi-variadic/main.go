package main

import "fmt"

func main() {
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)
	slice := []int{20, 20, 20, 20, 20}
	total = sumAll(slice...)
	fmt.Println(total)
}
func sumAll(number ...int) int {
	total := 0

	for _, value := range number {
		total += value
	}
	return total
}
