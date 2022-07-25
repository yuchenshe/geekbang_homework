package main

import "fmt"

func main() {
	var myArray [3]int
	for i := 0; i < 3; i++ {
		myArray[i] = i
	}
	fmt.Printf("myArray %+v\n", myArray)

	Array := [5]int{1, 2, 3, 4, 5}
	mySlice := Array[1:]
	fmt.Printf("mySlice %+v\n", mySlice)
}
