package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type myCounter interface {
	Inc()
	Load() int64
}

//普通版
type myCommandCounter struct {
	counter int64
}

func (c myCommandCounter) Inc() {
	c.counter++
}
func (c myCommandCounter) Load() int64 {
	return c.counter
}

// 互斥锁版
type myMutexCounter struct {
	counter int64
	lock    sync.Mutex
}

func (m *myMutexCounter) Inc() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

func (m *myMutexCounter) Load() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

//原子操作版
type myAtomicCounter struct {
	counter int64
}

func (a *myAtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}
func (a *myAtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func test(c myCounter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}

func main() {
	//线程不安全
	c1 := &myCommandCounter{} //0 1.207722ms
	test(c1)
	//线程安全
	c2 := &myMutexCounter{} //1000 1.132681ms
	test(c2)
	//线程安全且效率更高
	c3 := &myAtomicCounter{} //1000 523.061µs
	test(c3)
}
