package main

import "fmt"

func make_map() {
	fmt.Println("make_map()")
	dict := make(map[string]int)
	dict1 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	dict2 := map[int][]string{}

	dict["R"] = 11
	dict["L"] = 22

	for index, value := range dict {
		fmt.Println(index, value)
	}

	for index, value := range dict1 {
		fmt.Println(index, value)
	}

	t := []string{"AA", "AAA"}
	dict2[1] = t
	t = []string{"BB", "BBB"}
	dict2[2] = t

	for index, value := range dict2 {
		fmt.Println(index, value)
	}

	dict2[1] = append(dict2[1], "AAAA")
	dict2[2] = append(dict2[2], "BBBB")

	for index, value := range dict2 {
		fmt.Println(index, value)
	}
}

func nil_map() {
	fmt.Println("nil_map()")
	var colors map[string]int
	// colors["Red"] = "#da1337"    // runtime error (nil map 에 key/value 저장 안됨)

	m := make(map[string]int)
	m["Red"] = 11
	m["Blue"] = 22

	colors = m
	for key, value := range colors {
		fmt.Println(key, value)
	}
}

func select_map() {
	fmt.Println("select_map()")
	colors := make(map[string]int)
	colors["Red"] = 11
	colors["Blue"] = 22

	value, exists := colors["Blue"]
	if exists {
		fmt.Println(value)
	}
}

func check_exist() {
	fmt.Println("check_exist()")
	colors := make(map[string]int)
	colors["Red"] = 11
	colors["Blue"] = 22

	value := colors["Red"]
	if value != 0 {
		fmt.Println("Red exists")
	}

	value1, exists := colors["Blue"]
	fmt.Println(value1, exists)

	value1, exists = colors["Black"]
	fmt.Println(value1, exists)
}

func delete_elem() {
	fmt.Println("delete_elem()")

	colors := make(map[string]int)
	colors["Red"] = 11
	colors["Blue"] = 22

	for key, value := range colors {
		fmt.Println(key, value)
	}

	delete(colors, "Blue")

	for key, value := range colors {
		fmt.Println(key, value)
	}
}

func func_map() {
	fmt.Println("func_map()")

	colors := map[string]string{
		"AliceBlue":   "#11111",
		"Coral":       "#22222",
		"DarkGray":    "#33333",
		"ForestGreen": "#44444",
	}

	for key, value := range colors {
		fmt.Printf("key : %s value : %s\n", key, value)
	}

	removeColor(colors, "Coral")

	fmt.Println("======")
	for key, value := range colors {
		fmt.Printf("key : %s value : %s\n", key, value)
	}
}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}

func main() {
	make_map()
	nil_map()
	select_map()
	check_exist()
	delete_elem()
	func_map()
}
