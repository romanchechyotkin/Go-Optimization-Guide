package perf

import (
	"sync"
	"testing"
)

type SharedCounterBad struct {
	a int64
	b int64
}

type SharedCounterGood struct {
	a int64
	_ [56]byte // Padding to prevent a and b from sharing a cache line
	b int64
}

func BenchmarkFalseSharing(b *testing.B) {
	var c SharedCounterBad
	var wg sync.WaitGroup

	for b.Loop() {
		wg.Add(2)
		go func() {
			for i := 0; i < 1_000_000; i++ {
				c.a++
			}
			wg.Done()
		}()
		go func() {
			for i := 0; i < 1_000_000; i++ {
				c.b++
			}
			wg.Done()
		}()
		wg.Wait()
	}
}

func BenchmarkNoFalseSharing(b *testing.B) {
	var c SharedCounterGood
	var wg sync.WaitGroup

	for b.Loop() {
		wg.Add(2)
		go func() {
			for i := 0; i < 1_000_000; i++ {
				c.a++
			}
			wg.Done()
		}()
		go func() {
			for i := 0; i < 1_000_000; i++ {
				c.b++
			}
			wg.Done()
		}()
		wg.Wait()
	}
}
