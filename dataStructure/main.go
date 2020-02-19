package main

import (
	"fmt"

	"github.com/Westringliu/learnGo/dataStructure/ArrayList"
	"github.com/Westringliu/learnGo/dataStructure/StackArray"
)

func main1() {
	//list := ArrayList.NewArrayList()
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	fmt.Println(list)
	list.Delete(2)
	list.Delete(9)
	fmt.Println(list)
}
func main2() { //测试插入
	//list := ArrayList.NewArrayList()
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	fmt.Println(list)
	list.Delete(2)
	list.Delete(9)
	list.Insert(1, "a1")
	for i := 0; i < 10; i++ {
		list.Insert(i+2, "haha")
	}
	fmt.Println(list)
}
func main3() { //测试Iterator
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	list.Append("d")
	list.Append("e")
	list.Append("f")
	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next()
		if item == "d" {
			it.Remove()
		}
		fmt.Println(item)
	}
	fmt.Println(list)
}
func main4() { //测试StackArray
	// var mys StackArray.Stack = StackArray.NewStackArray()
	// mys.push("a1") //共有方法一定要大写，否则访问不了！！！！！！
	myStack := StackArray.NewStackArray()

	for i := 0; i < 12; i++ {
		err := myStack.Push(i)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("stack size:", myStack.Size())
	for i := 0; i < 11; i++ {
		ele, err := myStack.Pop()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(ele)
	}
	fmt.Println("stack size:", myStack.Size())
}

//栈模拟递归
func main5() {
	var mystack StackArray.Stack = StackArray.NewStackArray()
	mystack.Push(5)
	last := 0
	for !mystack.IsEmpty() {
		data, _ := mystack.Pop()
		if data == 0 {
			last += 0
		} else {
			last += data.(int)
			mystack.Push(data.(int) - 1)
		}
	}
	fmt.Println(last)
}

//纯递归实现斐波那契数列
func FAB(num int) int {
	//fmt.Println(num)
	if num == 2 {
		return 2
	} else if num == 1 {
		return 1
	} else {
		return FAB(num-1) + FAB(num-2)
	}
}
func main6() {
	for i := 1; i < 10; i++ {
		fmt.Println(FAB(i))
	}
}

//栈实现斐波那契数列
func FAB2(num int) int {
	var mystack StackArray.Stack = StackArray.NewStackArray()
	mystack.Push(num)
	last := 0
	for !mystack.IsEmpty() {
		data, _ := mystack.Pop()
		if data == 1 || data == 2 {
			last += 1
			//fmt.Println(data)
		} else {
			mystack.Push(data.(int) - 1)
			mystack.Push(data.(int) - 2)
		}
	}
	return last
	//fmt.Println(last)
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(FAB2(i))
	}
}
