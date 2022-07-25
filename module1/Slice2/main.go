package main

import "fmt"

func main() {
	a := []int{}
	b := []int{1, 2, 3}
	a = append(b, 1)
	c := a
	d := b[0]

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)

	updateSliceValue2(b)
}

//修改切片的值，在 range方法中修改值，只是修改了value，但其实切片本身内存的值并没有改动

func updateSliceValue(slice []int) {
	//这个循环就可以看出尽管修改了value，但是原切片并没有改变，因为这里的value作为一个新变量是一个指向的是一个新的内存地址，并不是指向的切片的内存地址
	for _, value := range slice {
		value *= 2

	}
	fmt.Println(slice)

}

//正确的修改切片值的函数
func updateSliceValue2(slice []int) {
	for index, _ := range slice {
		fmt.Println(index)
		slice[index] *= 2
	}
	fmt.Println(slice)
}
