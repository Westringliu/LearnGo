package main

import (
	"fmt"
	"os"
)

//空接口的使用--实现日志系统
type logSystem interface {
	consol(interface{})
	wfile(interface{})
}
type logSystem2 interface {
	consol2(interface{})
	wfile2(interface{})
}
type localLog struct{}
type myStruct struct {
	name string
	age  int
}

func (l localLog) consol(log interface{}) {
	loginfo := log.(string) + "\n"
	os.Stdout.WriteString(loginfo)
}
func (l localLog) wfile(log interface{}) {
	logMsg := log.(string) + "\n"
	_, err := os.Stat("interface.log")
	if err != nil {
		fmt.Println("文件不存在!需要创建！")
		// fmt.Println(err)
		file, err := os.OpenFile("interface.log", os.O_APPEND|os.O_CREATE|os.O_EXCL, 0664) //|os.O_CREATE|os.O_EXCL
		if err != nil {
			// fmt.Println(err)
			fmt.Printf("创建打开文件，报错：%s", err)
		}
		defer file.Close()
		if _, err := file.WriteString(logMsg); err == nil {
			fmt.Println("写入文件成功！")
		} else {
			// fmt.Println("文件写入报错！")
			// fmt.Println(err)
			fmt.Printf("文件写入报错！，报错信息：%s", err)
		}
	} else {
		fmt.Println("文件存在！")
		// fmt.Println(fileStat)
		file, err := os.OpenFile("interface.log", os.O_RDWR, 0666) //|os.O_CREATE|os.O_EXCL
		if err != nil {
			fmt.Printf("打开文件错误，报错：%s", err)
		}
		defer file.Close()
		if _, err := file.WriteString(logMsg); err == nil {
			fmt.Println("写入文件成功！")
		} else {
			fmt.Println("文件写入报错！")
			fmt.Println(err)
		}
	}
}

func (l localLog) consol2(log interface{}) {
	//自定义结构体作为空接口传入没有意义，底层并不知道也不应该知道上层数据结构的实现
	// loginfo := log.(string) + "\n"
	// logint := log.(int)
	// os.Stdout.WriteString(loginfo + string(logint))
	//不过依然可以实现自定义结构体作为空接口传入，可见其也是有空接口的
	fmt.Println(log)
}
func (l localLog) wfile2(log interface{}) {
	fmt.Println(log)
	// logMsg := log.(string) + "\n"
	// _, err := os.Stat("interface.log")
	// if err != nil {
	// 	fmt.Println("文件不存在!需要创建！")
	// 	// fmt.Println(err)
	// 	file, err := os.OpenFile("interface.log", os.O_APPEND|os.O_CREATE|os.O_EXCL, 0664) //|os.O_CREATE|os.O_EXCL
	// 	if err != nil {
	// 		// fmt.Println(err)
	// 		fmt.Printf("创建打开文件，报错：%s", err)
	// 	}
	// 	defer file.Close()
	// 	if _, err := file.WriteString(logMsg); err == nil {
	// 		fmt.Println("写入文件成功！")
	// 	} else {
	// 		// fmt.Println("文件写入报错！")
	// 		// fmt.Println(err)
	// 		fmt.Printf("文件写入报错！，报错信息：%s", err)
	// 	}
	// } else {
	// 	fmt.Println("文件存在！")
	// 	// fmt.Println(fileStat)
	// 	file, err := os.OpenFile("interface.log", os.O_RDWR, 0666) //|os.O_CREATE|os.O_EXCL
	// 	if err != nil {
	// 		fmt.Printf("打开文件错误，报错：%s", err)
	// 	}
	// 	defer file.Close()
	// 	if _, err := file.WriteString(logMsg); err == nil {
	// 		fmt.Println("写入文件成功！")
	// 	} else {
	// 		fmt.Println("文件写入报错！")
	// 		fmt.Println(err)
	// 	}
	// }
}

func logger(log string) {
	// log := log
	var logger logSystem
	var application = localLog{}
	logger = application
	logger.consol(log)
	logger.wfile(log)

}

func mylogger(mys myStruct) {
	var logger logSystem2
	var app = localLog{}
	logger = app
	logger.consol2(mys)
	logger.wfile2(mys)
}

func main() {
	var word string = "阿里巴巴与四十大盗111!"
	logger(word)
	var mys = myStruct{
		name: "xiaoming",
		age:  15,
	}
	mylogger(mys)
}
