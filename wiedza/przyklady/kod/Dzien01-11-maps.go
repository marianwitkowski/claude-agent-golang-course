package main

import (
	"fmt"
)

func findKey(key string, dict map[string]int) bool {
	for k := range dict {
		if k == key {
			return true
		}
	}
	return false
}

func main() {
	var regPlates = map[string]string{"WI": "Warszawa", "GD": "Gdańsk", "PO": "Poznań"}
	emptyMap := map[string]int{}

	a := make(map[string]int)
	a["key1"] = 10
	a["key2"] = 20

	fmt.Println(regPlates)
	fmt.Println(emptyMap)
	fmt.Println(a)
	if findKey("key3", a) {
		fmt.Println(a["key3"])
	} else {
		fmt.Println("brak klucza")
	}

}
