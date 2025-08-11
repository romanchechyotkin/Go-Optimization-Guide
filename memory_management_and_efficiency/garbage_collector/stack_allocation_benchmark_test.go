package perf

import "testing"

type Data struct {
	A, B, C int
}

func StackAlloc() Data {
	return Data{1, 2, 3} // stays on stack
}

func HeapAlloc() *Data {
	return &Data{1, 2, 3} // escapes to heap
}

func BenchmarkStackAlloc(b *testing.B) {
	for b.Loop() {
		_ = StackAlloc()
	}
}

func BenchmarkHeapAlloc(b *testing.B) {
	for b.Loop() {
		_ = HeapAlloc()
	}
}

var sink *Data

func HeapAllocEscape() {
	d := &Data{1, 2, 3}
	sink = d // d escapes to heap
}

func BenchmarkHeapAllocEscape(b *testing.B) {
	for b.Loop() {
		HeapAllocEscape()
	}
}
