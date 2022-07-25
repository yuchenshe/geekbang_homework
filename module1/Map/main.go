package main

import "fmt"

func main() {
	myMap := map[string]int{
		"syc":  1,
		"yshe": 2,
	}
	for k, v := range myMap {
		if k == "syc" {
			myMap[k] = v * 2
		}
	}
	fmt.Println(myMap)
	myFuncMap := map[string]func() int{"funcA": func() int { return 1 }}
	fmt.Println(myFuncMap)
	f := myFuncMap["funcA"]
	fmt.Println(f())
}
