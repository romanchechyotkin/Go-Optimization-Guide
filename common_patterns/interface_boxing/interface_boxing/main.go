package main

import (
	"log"
	"unsafe"
)

type Shape interface {
	Area() float64
}

type Square struct {
	Size float64
}

func (s Square) Area() float64 { return s.Size * s.Size }

// Pay attention to this code!
// In this example, even though shapes is a slice of interfaces, each Square value is copied into an interface when appended to shapes.
// If Square were a large struct, this would introduce 1000 allocations and large memory copying.
func main() {
	var shapes []Shape
	for i := 0; i < 1000; i++ {
		s := Square{Size: float64(i)}
		shapes = append(shapes, s) // boxing occurs here
	}
	log.Printf("shapes len=%d; cap=%d; size=%d", len(shapes), cap(shapes), unsafe.Sizeof(shapes))

	shapes = []Shape{}
	for i := 0; i < 1000; i++ {
		s := Square{Size: float64(i)}
		shapes = append(shapes, &s) // boxing occurs here
	}
	log.Printf("shapes len=%d; cap=%d; size=%d", len(shapes), cap(shapes), unsafe.Sizeof(shapes))
}
