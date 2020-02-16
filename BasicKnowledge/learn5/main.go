package main

import "fmt"

func adder(x int) func(int) int {
	return func(y int) int { //闭包 = 函数 + 引用环境
		x += y
		return x
	}
}

func main() {
	fmt.Println("<<<<<<<<<<<<测试闭包函数")
	{
		var f = adder(10) //f就是adder返回的闭包函数，因此传入f的参数就是y，x则作为上下文在闭包函数的生命周期内始终被记录
		fmt.Println(f(10))
		fmt.Println(f(20))
		fmt.Println(f(30))
	} //主要用于解释生命周期和闭包的关系
	var f = adder(10) //f就是adder返回的闭包函数，因此传入f的参数就是y，x则作为上下文在闭包函数的生命周期内始终被记录
	fmt.Println(f(10))
	fmt.Println(f(20))
	fmt.Println(f(30))
	fmt.Println("<<<<<<<<<<<<测试defer延迟执行")
	fmt.Println("begin") //猜测defer的实现是基于栈，主程执行时遇到defer就讲后面的操作压栈，主程退出前依次出栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}
