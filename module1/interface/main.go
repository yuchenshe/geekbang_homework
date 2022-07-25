package main

import "fmt"

type jiekou interface {
	getName() string
}

type human struct {
	first_name, last_name string
}

type car struct {
	vendor, model string
}

type plane struct {
	vendor, model string
}

func (h human) getName() string {
	return h.first_name + "-" + h.last_name
}

func (c car) getName() string {
	return c.vendor + "-" + c.model
}

func (p plane) getName() string {
	return p.vendor + "," + p.model
}

func main() {
	h := new(human)
	h.first_name = "yuchen"
	h.last_name = "she"
	c := new(car)
	c.vendor = "bmw"
	c.model = "330li"

	myslice := []jiekou{}

	myslice = append(myslice, h, c)

	for _, value := range myslice {
		fmt.Println(value.getName())
	}

	p := new(plane)
	p.vendor = "keji"
	p.model = "donghang"
	fmt.Println(p.getName())

}
