package main

import (
	"sync"
	"sync/atomic"
)

type LazySingleton struct{}

var lazySingleton *LazySingleton

// 该版本存在数据竞争（data race）问题
// A data race occurs when two goroutines access the same variable concurrently and at least one of the accesses is a write.
// 至少有两个 goroutine 同时去访问一个变量，而这两个 goroutine 中至少有一个会写这个变量。
var mutex sync.Mutex

func getLazyInstanceByMutex() *LazySingleton {
	if lazySingleton == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if lazySingleton == nil {
			lazySingleton = &LazySingleton{}
		}
	}
	return lazySingleton
}

// 检测数据竞争
// 执行go run --race Singleton_lazy.go
func main() {
	for i := 0; i < 5; i++ {
		go func() {
			getLazyInstanceByMutex()
		}()
	}

}

var rwmutex sync.RWMutex

func getLazySingletonByRWMutex() *LazySingleton {
	rwmutex.RLock()
	rwmutex.RUnlock()
	if lazySingleton == nil {
		rwmutex.Lock()
		defer rwmutex.Unlock()
		if lazySingleton == nil {
			lazySingleton = &LazySingleton{}
		}
	}
	return lazySingleton
}

var (
	mutexWithAtomic sync.Mutex
	flag            uint32
)

func getLazySingletonByAtomic() *LazySingleton {
	if atomic.LoadUint32(&flag) == 0 {
		mutex.Lock()
		defer mutex.Unlock()
		if atomic.LoadUint32(&flag) == 0 {
			lazySingleton = &LazySingleton{}
			defer atomic.StoreUint32(&flag, 1)
		}
	}
	return lazySingleton
}

var once sync.Once

func getLazySingletonByOnce() *LazySingleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &LazySingleton{}
		})
	}
	return lazySingleton
}
