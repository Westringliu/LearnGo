package ArrayList

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int
	Get(index int) (interface{}, error)
	Set(index int, newval interface{}) error
	Insert(index int, newval interface{}) error
	Append(newval interface{})
	Clear()
	Delete(index int) error
	String() string
	Iterator() Iterator
}

type ArrayList struct {
	dataStore []interface{}
	theSize   int
}

// NewArrayList fasfdas
func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.dataStore = make([]interface{}, 0, 10)
	list.theSize = 0
	return list
}

func (list *ArrayList) Size() int {
	return list.theSize
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.theSize {
		return nil, errors.New("array over flow")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Set(index int, newval interface{}) error {
	if index < 0 || index >= list.theSize {
		return errors.New("array over flow")
	}
	list.dataStore[index] = newval
	return nil
}

func (list *ArrayList) checkisFull() {
	if list.theSize == cap(list.dataStore) {
		newdataStore := make([]interface{}, 2*list.theSize, 2*list.theSize)
		copy(newdataStore, list.dataStore)
		list.dataStore = newdataStore
	}
}
func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index < 0 || index >= list.theSize {
		return errors.New("array over flow")
	}
	list.checkisFull()
	list.dataStore = list.dataStore[:list.theSize+1] //插入的时候需要将内存移动一位
	for i := list.theSize; i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newval
	list.theSize++
	return nil
}
func (list *ArrayList) Append(newval interface{}) {
	list.dataStore = append(list.dataStore, newval)
	list.theSize++
}
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.theSize = 0
}
func (list *ArrayList) Delete(index int) error {
	if index < 0 || index >= list.theSize {
		return errors.New("array over flow")
	}
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...) //不加...输出是[1 2 [4 5 6]]
	list.theSize--
	return nil
}
func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}
