package main

import (
	"log"
	"sync"
)

type Counter struct {
	a int64
	b int64
}

type Counter2 struct {
	a int64
	_ [56]byte
	b int64
}

func main() {
	var c Counter
	var wg sync.WaitGroup

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

	log.Printf("%+v", c)

	var c2 Counter2
	wg.Add(2)

	go func() {
		for i := 0; i < 1_000_000; i++ {
			c2.a++
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1_000_000; i++ {
			c2.b++
		}
		wg.Done()
	}()

	wg.Wait()

	log.Printf("%+v", c2)
}
