package main

import "fmt"

func main() {
	age := 44
	if age <= 12 {
		fmt.Println("children")
	} else if age > 12 && age <= 18 {
		fmt.Println("boy")
	} else if age > 18 && age <= 50 {
		fmt.Println("man")
	} else {
		fmt.Println("old")
	}
	var a = [3]string{"hello", "byebye", "haha"} //类型是[3]int
	var b [4]int                                 //类型是[4]int  所以不可以a = b
	var c = [2]int{0, 1}
	var d = [...]string{"a", "b", "c"}
	for i := 0; i < len(c); i++ {
		fmt.Println("a:", a[i], "b:", b[i], "c:", c[i], "d:", d[i])
	}
	for _, va := range a {
		fmt.Print(va, ",")
	}
	fmt.Println()
	fmt.Println(d)
	var e = [2][3]string{
		{"nihao", "hello", "hi"},
		{"zaijain", "byebye", "bye"},
	}
	fmt.Println(e)
	array := [...]int{1, 3, 5, 7, 8}
	tmpmap := map[int]int{}
	for i := 0; i < len(array); i++ {
		value, ok := tmpmap[8-array[i]]
		if ok {
			fmt.Println(value, i)
		} else {
			tmpmap[array[i]] = i
		}
	}
	for k, v := range tmpmap {
		fmt.Println(k, v)
	}
}
