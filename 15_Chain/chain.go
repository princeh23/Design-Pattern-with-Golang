package main

import "fmt"

var status int8 = 0

type HandlerFunc func()

type HandlersChain []HandlerFunc

type RouterGroup struct {
	Handlers HandlersChain
	index    int8
}

func (group *RouterGroup) Use(middleware ...HandlerFunc) {
	group.Handlers = append(group.Handlers, middleware...)
}

func (group *RouterGroup) Next() {
	fmt.Println(group.index)
	for group.index < int8(len(group.Handlers)) {
		group.Handlers[group.index]()
		group.index++
	}
}

func middleware1() {
	fmt.Println("全局中间件1执行完毕")
}

func middleware2() {
	fmt.Println("全局中间件2执行完毕")
}

func main() {
	r := &RouterGroup{}
	r.Use(middleware1, middleware2)
	r.Next()
}
