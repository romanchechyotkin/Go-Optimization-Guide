package perf

import "testing"

func BenchmarkCopy(b *testing.B) {
	data := make([]byte, 64*1024)
	for b.Loop() {
		buf := make([]byte, len(data))
		copy(buf, data)
	}
}

func BenchmarkSlice(b *testing.B) {
	data := make([]byte, 64*1024)
	for b.Loop() {
		_ = data[:]
	}
}
