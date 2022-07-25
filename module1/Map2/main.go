package main

import "fmt"

func main() {
	myMap := map[string]int{"k1": 1, "k2": 2}
	for k, v := range myMap {
		fmt.Println(k, v)
	}

	fmt.Println(myMap)
	//getMapValue(myMap, "k")

	//delete(myMap, "k2")
	//fmt.Println(myMap)
}

func getMapValue(myMap map[string]int, k string) int {
	value, exists := myMap[k]
	if !exists {
		fmt.Println(k, " is not in this map")
	}
	return value
}
