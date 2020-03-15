package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/Westringliu/learnGo/dataStructure/StackArray"
)

func GetALL(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("can not read this dir")
	}
	for _, fi := range read {
		if fi.IsDir() {
			fulldir := path + "/" + fi.Name()
			files = append(files, fulldir)
			files, _ = GetALL(fulldir, files)
		} else {
			fulldir := path + "/" + fi.Name()
			files = append(files, fulldir)
		}
	}
	return files, nil
}

func GetALLX(path string, files []string, level int) ([]string, error) {
	var levelstr string
	tmpLevel := level
	//fmt.Println(level)
	if tmpLevel == 0 {
		levelstr = ""
	}
	for ; tmpLevel > 0; tmpLevel-- {
		levelstr += "|--"
	}
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("can not read this dir")
	}
	for _, fi := range read {
		if fi.IsDir() {
			fulldir := path + "/" + fi.Name()
			files = append(files, levelstr+fulldir)
			files, _ = GetALLX(fulldir, files, level+1)
		} else {
			fulldir := path + "/" + fi.Name()
			files = append(files, levelstr+fulldir)
		}
	}
	return files, nil
}

//递归实现文件夹遍历
func main1() {
	path := "/Users/westring/go/src/github.com/Westringliu/learnGo"
	files := []string{}
	files, _ = GetALL(path, files)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

//递归实现文件夹遍历（带层级现实）
func main2() {
	path := "/Users/westring/go/src/github.com/Westringliu/learnGo"
	files := []string{}
	files, _ = GetALLX(path, files, 0)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

//栈实现文件夹遍历
func main() {
	path := "/Users/westring/go/src/github.com/Westringliu/learnGo"
	files := []string{}

	mystack := StackArray.NewStackArray()
	mystack.Push(path)
	for !mystack.IsEmpty() {
		path := mystack.Pop().(string)
		//files = append(files, path)
		read, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Println(err)
		}
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path + "/" + fi.Name()
				files = append(files, fulldir)
				mystack.Push(fulldir)
			} else {
				fulldir := path + "/" + fi.Name()
				files = append(files, fulldir)
			}
		}
	}
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}
