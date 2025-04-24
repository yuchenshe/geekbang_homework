package main

import (
	"fmt"
)

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// 提前退出的标志
		swapped := false

		// 每一轮把最大的“冒”到后面
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 交换相邻元素
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		// 如果一轮都没交换，说明已经排好序了
		if !swapped {
			break
		}
	}
}

func main() {
	arr := []int{5, 3, 8, 4, 2}
	fmt.Println("排序前:", arr)

	bubbleSort(arr)

	fmt.Println("排序后:", arr)
}
