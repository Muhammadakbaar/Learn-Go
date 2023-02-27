package main

import "fmt"

func main() {

	var point1 = 8

	if point1 == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point1 > 5 {
		fmt.Println("lulus")
	} else if point1 == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point1)
	}

	var point2 = 6

	switch point2 {
	case 8:
		fmt.Println("Perfect")
	case 7, 6, 5, 4:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not bad")

	}

	var point3 = 3

	switch point3 {
	case 8:
		fmt.Println("Perfect")
	case 7, 6, 5, 4:
		fmt.Println("Awesome")
	default:
		{
			fmt.Println("Not bad")
			fmt.Println("You can do better")
		}

	}
	var point4 = 6

	switch {
	case point4 == 8:
		fmt.Println("perfect")
	case (point4 < 8) && (point4 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}

	}
	switch {
	case point4 == 8:
		fmt.Println("perfect")
	case (point4 < 8) && (point4 > 3):
		fmt.Println("awesome")
		fallthrough
	case point4 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

}
