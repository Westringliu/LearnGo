package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello", i)
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main() {
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功")
}

/*
func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}
*/
