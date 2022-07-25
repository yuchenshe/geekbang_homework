package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A string
	B string
}

func (t T) String() string {
	return t.A + "yshe"
}

func main() {
	myMap := make(map[string]string)
	myMap["a"] = "she"
	t := reflect.TypeOf(myMap)
	fmt.Println(t)
	v := reflect.ValueOf(myMap)
	fmt.Println(v)

	myStruct := T{A: "a", B: "b"}
	v1 := reflect.ValueOf(myStruct)

	for i := 0; i < v1.NumField(); i++ {
		fmt.Println(v1.Field(i))
	}

	for i := 0; i < v1.NumMethod(); i++ {
		fmt.Printf("Method %d: %v\n", i, v1.Method(i))
	}

	fmt.Println(v1.Method(0).Call(nil))

	func(x, y int) {
		fmt.Println(x + y)
	}(1, 2)

}
