package perf

import "testing"

type PoorlyAligned struct {
	flag  bool  // 1 byte + 7 byte for alignment
	count int64 // 8 byte                        => 24 bytes in total
	id    byte  // 1 byte + 7 byte for alignment
} // in toatl 14 bytes for alignment

type WellAligned struct {
	count int64 // 8 bytes
	flag  bool  // 1 byte } => 8 bytes => 16 bytes in total
	id    byte  // 1 byte }
} // in total 6 bytes for alignment

func BenchmarkPoorlyAligned(b *testing.B) {
	for b.Loop() {
		var items = make([]PoorlyAligned, 10_000_000)
		for j := range items {
			items[j].count = int64(j)
		}
	}
}

func BenchmarkWellAligned(b *testing.B) {
	for b.Loop() {
		var items = make([]WellAligned, 10_000_000)
		for j := range items {
			items[j].count = int64(j)
		}
	}
}
