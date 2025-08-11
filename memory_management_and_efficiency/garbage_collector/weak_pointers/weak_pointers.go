package main

import (
	"fmt"
	"runtime"
	"weak"
)

type Data struct {
	Value string
}

func main() {
	data := &Data{Value: "Important"}
	wp := weak.Make(data) // create weak pointer

	fmt.Println("Original:", wp.Value().Value)

	data = nil // remove strong reference
	runtime.GC()

	if v := wp.Value(); v != nil {
		fmt.Println("Still alive:", v.Value)
	} else {
		fmt.Println("Data has been collected")
	}
}
