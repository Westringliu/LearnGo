package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
	Remove()
	GetIndex() int
}

type Itrable interface {
	Iterator() Iterator
}

type ArrayListIterator struct {
	list         *ArrayList
	currentindex int
}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator)
	it.currentindex = 0
	it.list = list
	return it
}

func (it *ArrayListIterator) HasNext() bool {
	return it.currentindex < it.list.theSize
}
func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("no next")
	}
	value, err := it.list.Get(it.currentindex)
	it.currentindex++
	return value, err
}
func (it *ArrayListIterator) Remove() {
	it.currentindex--
	it.list.Delete(it.currentindex)
}
func (it *ArrayListIterator) GetIndex() int {
	return it.currentindex
}
