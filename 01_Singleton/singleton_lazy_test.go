package main

import (
	"testing"
)

var N = 10000

// go test -bench . -benchtime=5s .

func Benchmark_getLazyInstanceByMutex(b *testing.B) {
	for i := 0; i < N; i++ {
		getLazyInstanceByMutex()
	}
}

func Benchmark_getLazyInstanceByRWMutex(b *testing.B) {
	for i := 0; i < N; i++ {
		getLazySingletonByRWMutex()
	}
}

func Benchmark_getLazyInstanceByAtomic(b *testing.B) {
	for i := 0; i < N; i++ {
		getLazySingletonByAtomic()
	}
}

func Benchmark_getLazyInstanceByOnce(b *testing.B) {
	for i := 0; i < N; i++ {
		getLazySingletonByOnce()
	}
}
