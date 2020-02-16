package main

import (
	"fmt"
	"sync"
	"time"
)

var x int
var wg sync.WaitGroup

func worker(id int, jobs <-chan int, resault chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d strat job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		resault <- j * 2
	}
}

func Add() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		x += 1
	}
}

func main() {
	jobs := make(chan int, 100)
	resault := make(chan int, 100)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, resault)
	}
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	wg.Add(2)
	go Add()
	go Add()
	wg.Wait()
	fmt.Println(x)
	//close(resault)
	// for i := range resault {
	// 	fmt.Println(i)
	// }
	for i := 1; i <= 5; i++ {
		<-resault
	}
	close(jobs)
	close(resault) //可有可无，chan可以自己关闭
}
