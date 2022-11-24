package main

import (
	"fmt"
)

// Iterator 迭代器接口
type Iterator interface {
	HasNext() bool
	Next()
	// 获取当前元素，由于 Go 1.15 中还没有泛型，所以我们直接返回 interface{}
	CurrentItem() interface{}
}

type ArrayInt []int

func (a ArrayInt) Iterator() Iterator {
	return &ArrayIntIterator{
		arrayInt: a,
		index:    0,
	}
}

type ArrayIntIterator struct {
	arrayInt ArrayInt
	index    int
}

func (iter *ArrayIntIterator) HasNext() bool {
	return iter.index < len(iter.arrayInt)
}

func (iter *ArrayIntIterator) Next() {
	iter.index++
}

func (iter *ArrayIntIterator) CurrentItem() interface{} {
	return iter.arrayInt[iter.index]
}

func main() {
	array := ArrayInt{1, 3, 5, 7, 9}
	it := array.Iterator()
	for it.HasNext() {
		fmt.Println(it.CurrentItem())
		it.Next()
	}
}
