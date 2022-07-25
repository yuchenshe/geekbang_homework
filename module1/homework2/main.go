package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)

	go producer(ch)
	//消费者
	func() {
		for i := range ch {
			fmt.Println("receive:", i)
		}

	}()
}

//生产者
func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}

	close(ch)
}
