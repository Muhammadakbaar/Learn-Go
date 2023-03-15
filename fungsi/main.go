package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	var names = []string{"John", "Wick"}
	printMessage("Hello", names)
	coba("mneteng", map[string]string{"menteng": "menteng"})

	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number :", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number :", randomValue)
	randomValue = randomWithRange(-2, 2)
	fmt.Println("Random number :", randomValue)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
func coba(coba1 string, lana map[string]string) {
	fmt.Println(coba1)
}

func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}
