package main

import "fmt"

func main() {
	var name string = "xiaoming"
	var age uint = 18
	var name1, age1 = "laowang", 20
	name2 := "lizi"
	age2 := 22
	const (
		pi         = 3.14
		e  float32 = 2.71
		f
	)
	fmt.Println("你好啊")
	fmt.Println(name, age, name1, age1, name2, age2)
	fmt.Println(pi, e, f)
}
