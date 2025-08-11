package main

import (
	"log"
)

func main() {
	s := make([]int, 0)
	for i := 0; i < 129; i++ {
		s = append(s, i)
		log.Printf("Len: %d, Cap: %d\n", len(s), cap(s))
	}

	// Inefficient
	var result []int
	for i := 0; i < 10000; i++ {
		result = append(result, i)
	}
	log.Printf("without prealloc esult slice len=%d; cap=%d\n", len(result), cap(result))

	// Efficient
	result = make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		result = append(result, i)
	}
	log.Printf("with prealloc result slice len=%d; cap=%d\n", len(result), cap(result))
}
