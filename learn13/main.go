package main

import (
	"fmt"
	"sync"
	"time"
)

/*
var wg sync.WaitGroup
var mutexLoc sync.Mutex //简单互斥锁
var rwloc sync.RWMutex  //读写锁分离，读锁之间不会阻塞，写锁会
var x int

func add() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mutexLoc.Lock() //多线程需要注意共享内存静态问题
		x++
		mutexLoc.Unlock()
	}
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
*/

var (
	x        int64
	wg       sync.WaitGroup
	mutexLoc sync.Mutex
	rwLoc    sync.RWMutex
)

func myWrite() {
	mutexLoc.Lock()
	//rwLoc.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	//rwLoc.Unlock()
	mutexLoc.Unlock()
	wg.Done()
}

func myRead() {
	mutexLoc.Lock()
	//rwLoc.RLock()	//大量读操作的情况下读写锁效率更高
	time.Sleep(time.Millisecond)
	//rwLoc.RUnlock()
	mutexLoc.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go myWrite()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go myRead()
	}
	go myRead()
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
