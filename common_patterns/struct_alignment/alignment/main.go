package main

import (
	"log"
	"unsafe"
)

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

func main() {
	p := PoorlyAligned{}
	log.Printf("size of empty poorly aligned struct value: %d", unsafe.Sizeof(p))

	pp := &p
	log.Printf("size of empty poorly aligned struct pointer: %d", unsafe.Sizeof(pp))

	pf := PoorlyAligned{
		flag:  true,
		count: 123124,
		id:    255,
	}

	log.Printf("size of filled poorly aligned struct value: %d", unsafe.Sizeof(pf))

	pp = &pf
	log.Printf("size of filled poorly aligned struct pointer: %d", unsafe.Sizeof(pp))

	w := WellAligned{}
	log.Printf("size of empty well aligned struct value: %d", unsafe.Sizeof(w))

	wp := &w
	log.Printf("size of empty well aligned struct pointer: %d", unsafe.Sizeof(wp))

	wf := WellAligned{
		flag:  true,
		count: 123124,
		id:    255,
	}

	log.Printf("size of filled well aligned struct value: %d", unsafe.Sizeof(wf))

	wp = &wf
	log.Printf("size of filled well aligned struct pointer: %d", unsafe.Sizeof(wp))
}
