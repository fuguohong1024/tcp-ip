package main

import "testing"

var (
	slice1 = []byte("my first slice")
	slice2 = []byte("second slice")
	slice3 = []byte("third slice")
	slice4 = []byte("fourth slice")
	slice5 = []byte("fifth slice")
)

var B []byte

// go test . -bench=. -benchmem

func BenchmarkConcatCopyPreAllocate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		B = concatCopyPreAllocate(slice1, slice2, slice3, slice4, slice5)
	}
}

func BenchmarkConcatAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		B = concatAppend(slice1, slice2, slice3, slice4, slice5)
	}
}

// 本文网址: https://golangnote.com/topic/188.html 转摘请注明来源
