package StackArray

import (
	"errors"
)

type Stack interface {
	IsFull() bool
	IsEmpty() bool
	Push(ele interface{}) error
	Pop() (interface{}, error)
	Size() int
	Clear()
}
type StackArray struct {
	stackStore  []interface{}
	capSize     int
	currentSize int
}

func NewStackArray() *StackArray {
	myStack := new(StackArray)
	myStack.stackStore = make([]interface{}, 0, 5000)
	myStack.capSize = 5000
	myStack.currentSize = 0
	return myStack
}

func (myStack *StackArray) IsFull() bool {
	return myStack.currentSize == myStack.capSize
}
func (myStack *StackArray) IsEmpty() bool {
	return myStack.currentSize == 0
}
func (myStack *StackArray) Push(ele interface{}) error {
	if myStack.IsFull() {
		return errors.New("stack is full")
	}
	myStack.currentSize++
	//以下这步很关键，不做会导致引发panic
	myStack.stackStore = myStack.stackStore[:myStack.currentSize] //插入的时候需要将内存移动一位
	myStack.stackStore[myStack.currentSize-1] = ele
	return nil
}
func (myStack *StackArray) Pop() interface{} {
	if myStack.IsEmpty() {
		return nil
	}
	myStack.currentSize--
	top := myStack.stackStore[myStack.currentSize]
	myStack.stackStore = myStack.stackStore[:myStack.currentSize]
	return top
}
func (myStack *StackArray) Size() int {
	return myStack.currentSize
}
func (myStack *StackArray) Clear() {
	myStack.stackStore = make([]interface{}, 0, 5000)
}
