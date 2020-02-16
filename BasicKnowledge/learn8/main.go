package main

import "fmt"

/*  //用interface实现的多态
type speaker interface {
	speak()
}

type cat struct{}
type dog struct{}
type people struct{}

func (c cat) speak() {
	fmt.Println("miaomiao")
}
func (d dog) speak() {
	fmt.Println("wangwang")
}
func (p people) speak() {
	fmt.Println("aa")
}
func da(x speaker) {
	x.speak()
}

func main() {
	var c cat
	var d dog
	var p people
	da(c)
	da(d)
	da(p)
}
*/

/*	接受者和接受者指针区别
type dog struct{}

type Mover interface {
	move() int
}

func (d *dog) move() int {
	fmt.Println("狗会动")
	return 1
}

func main() {
	var x Mover
	var wangcai = dog{} // 旺财是dog类型
	//x = wangcai         // 报错，x不可以接收dog类型
	var fugui = &dog{} // 富贵是*dog类型
	x = fugui          // x可以接收*dog类型
	x.move()
	fmt.Println(wangcai.move()) //按道理应该会报错啊，传指针类型才对，没搞明白为啥不报错
}
*/

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}
type gangan struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

func (d haier) dry() {
	fmt.Println("haier甩一甩")
}
func (g gangan) dry() {
	fmt.Println()
}

// 海尔洗衣机
type haier struct {
	dryer  //嵌入甩干器
	gangan //2号甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

func main() {
	var wm WashingMachine
	var d dryer = dryer{}
	var g gangan = gangan{}
	var h haier = haier{d, g}
	wm = h
	//m = h
	//自己定义了方法优先用自己的，没有再找嵌套类的,若多个嵌套类均实现了dry方法则需要指明，否则报错
	//h.dry()
	h.wash()
	h.dryer.dry()
	wm.dry()
	//fmt.Println(wm)
}

/*
//接口的嵌套
// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (c cat) move() {
	fmt.Println("猫会动")
}

func main() {
	var x animal
	x = cat{name: "花花"}
	x.move()
	x.say()
}
*/
