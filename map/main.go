package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["february"] = 40

	fmt.Println("januari ", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	var data map[string]int
	data = map[string]int{}
	data["one"] = 1
	fmt.Println("one", data["one"])

	var chicken1 = map[string]int{
		"segar":  50,
		"kurang": 40,
	}
	var chicken2 = map[int]int{
		1: 1,
		2: 2,
	}
	fmt.Println(chicken1["segar"], chicken1["kurang"], chicken2[1])

	var chicken3 = make(map[string]int)
	chicken3[""] = 1

	chicken3 = map[string]int{
		"data": 1,
	}

	fmt.Println(chicken3["data"])
	var chicken4 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}
	fmt.Println("=== map dengan for ===")
	for key, val := range chicken4 {
		fmt.Println(key, " \t:", val)
	}

	fmt.Println("=== delete item map ===")
	var chicken5 = map[string]int{

		"januari":  50,
		"februari": 40,
	}
	fmt.Println(len(chicken5))
	fmt.Println(chicken5)

	delete(chicken5, "januari")

	fmt.Println(len(chicken5))
	fmt.Println(chicken5)

	fmt.Println("=== deteksi item dengan key tertentu ===")
	var chicken6 = map[string]int{
		"januari":  50,
		"februari": 40,
	}
	var value, isExist = chicken6["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exist")
	}
	fmt.Println("=== kombinasi slice dan map ===")

	var chickens = []map[string]string{
		map[string]string{"nama": "chicken blue", "gender": "male"},
		map[string]string{"nama": "chicken red", "gender": "male"},
		map[string]string{"nama": "chicken yellow", "gender": "female"},
	}
	for _, chicken7 := range chickens {
		fmt.Println(chicken7["gender"], chicken7["nama"])
	}
}
