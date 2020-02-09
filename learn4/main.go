package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type calcu func(int, ...int) int

func sum(x int, y ...int) int {
	sum := x
	for _, value := range y {
		sum += value
	}
	return sum
}

func main() {
	fmt.Println("<<<<<<<<<<<<<切片的使用和排序")
	var a = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a[:])
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a[:])))
	fmt.Println(a)

	fmt.Println("<<<<<<<<<<<<<map的使用")
	mmp := make(map[string]int)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		mmp[key] = value
	}
	keys := make([]string, 0, 200)
	for key := range mmp {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, mmp[key])
	}

	fmt.Println("<<<<<<<<<<<<<func可变参数")
	s := sum(10, 20, 30, 40)
	fmt.Println(s)

	fmt.Println("<<<<<<<<<<<<<回调函数的定义和使用")
	var cal calcu
	cal = sum
	fmt.Println(cal(1, 2, 3, 4, 5))
}
