package main

import (
	"fmt"
	"log"
)

func main() {
	// Inefficient
	m0 := make(map[int]string)
	for i := 0; i < 10000; i++ {
		m0[i] = fmt.Sprintf("val-%d", i)
	}
	log.Println(len(m0))

	// Efficient
	m1 := make(map[int]string, 10000)
	for i := 0; i < 10000; i++ {
		m1[i] = fmt.Sprintf("val-%d", i)
	}
	log.Println(len(m1))
}
