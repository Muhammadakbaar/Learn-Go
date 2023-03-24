package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avg = calculate(numbers...)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)
	yourHobbies("Wick", "Sleeping ", "eating")

}

func calculate(number ...int) float64 {
	var total int = 0
	for _, number := range number {
		total += number
	}
	var avg = float64(total) / float64(len(number))
	return avg
}
func yourHobbies(name string, hobbies ...string) {
	var hobbiesasAsString = strings.Join(hobbies, ",")

	fmt.Printf("Hello my name is : %s\n", name)
	fmt.Printf("My hobbies are : %s\n", hobbiesasAsString)
}
