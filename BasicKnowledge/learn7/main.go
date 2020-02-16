package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]student, 3) //注意和map[string]*student结果比较
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = stu //stu的值不断变化，stu的地址不变
	}
	// for k, v := range m {	//直接赋值而不是地址的情况下无法直接寻址，参考https://blog.csdn.net/zhngcho/article/details/82424962
	// 	fmt.Println(k, "=>", v.name)
	// }
	fmt.Println(m)
}
