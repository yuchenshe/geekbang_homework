package main

import (
	"errors"
	"fmt"
)

func main() {
	var a string = "hello world!"
	fmt.Println(a)

	for i, e := range a {
		fmt.Println(i, string(e))
	}

	var myError error = errors.New("NotFound")

	fmt.Println(myError)

}
