package main

import "fmt"

func myFunc(benjin float32, lixi float32, nianshu int) float32 {
	if nianshu == 0 {
		return benjin
	}
	nianshu--
	benjin *= lixi + 1
	return myFunc(benjin, lixi, nianshu)
}

func main() {
	sum := myFunc(10000.0, 0.04, 20)
	fmt.Println(sum)
}
