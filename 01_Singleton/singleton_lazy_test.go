package main

import (
	"reflect"
	"testing"
)

var N = 10000

// go test -bench . -benchtime=5s .

func Benchmark_getLazyInstanceByMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
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

func Test_getLazySingletonByOnce(t *testing.T) {
	tests := []struct {
		name string
		want *LazySingleton
	}{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLazySingletonByOnce(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLazySingletonByOnce() = %v, want %v", got, tt.want)
			}
		})
	}
}
