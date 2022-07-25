package main

import "fmt"

func main() {
	myArray := [5]string{"I", "am", "stupid", "and", "weak"}

	for index, value := range myArray {

		switch value {
		case "stupid":
			myArray[index] = "smart"
		case "weak":
			myArray[index] = "strong"
		}

	}

	fmt.Println(myArray)

}
