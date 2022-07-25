package main

import "fmt"

var hisSlice []int

func main() {
	var slice []int
	if slice == nil {
		fmt.Println("slice为空")
	}

	//new一个新的切片
	mySlice1 := new([]int)
	fmt.Println("mySlice1:", mySlice1)

	//创建mySlice2，并且预设好长度和容量
	mySlice2 := make([]int, 10, 20)
	length := len(mySlice2)
	captain := cap(mySlice2)
	fmt.Println(length, captain)

	//切片没有内置的元素删除函数，所以要自己定义一个函数来删除其中的一个元素，要注意方括号中的范围是左包含，右不包含，且从0开始计数
	myArray := [7]int{1, 2, 3, 4, 5, 6, 7}
	mySlice := myArray[:]
	fmt.Printf("mySlice: %+v\n", mySlice)

	newSlice := deleteSliceItem(mySlice, 6)
	fmt.Println("newSlice:", newSlice)
}

func deleteSliceItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)

}
