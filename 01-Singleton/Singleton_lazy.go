package main

import "sync"

type LazySingleton struct{}

var lazySingleton *LazySingleton
var once sync.Once

func getLazyInstance() *LazySingleton {
	once.Do(func() {
		lazySingleton = &LazySingleton{}
	})
	return lazySingleton
}

func main() {
	getLazyInstance()
}
