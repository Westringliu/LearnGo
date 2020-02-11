package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

func Icon(name string) image.Image {
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}
*/

/*
//利用sync.Once实现单例模式
type singleton struct{}

var instance *singleton
var once sync.Once

func myGetInstance() *singleton {
	//once.Do(func() {	//没有加Once理应返回的对象地址值不同，现实却不如此，费解！！！
	instance = &singleton{} //该句有new操作
	//})
	return instance
}

func main() {
	s1 := myGetInstance()
	s2 := myGetInstance()
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)
}
*/

//内置map非线程安全
//var m = make(map[string]int)

//sync包提供了开箱即用的线程安全的map
var m = sync.Map{}

// func get(key string) int {
// 	return m[key]
// }
// func set(key string, val int) {
// 	m[key] = val
// }

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			//set(key, n)	//内置map线程不安全，并发量大运行时就会报错了
			//fmt.Printf("k = :%v,v := %v\n", key, get(key))
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k = :%v,v := %v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
