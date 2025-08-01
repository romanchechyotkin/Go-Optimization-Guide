package main

import (
	"fmt"
)

// Data is a struct with a large fixed-size array to simulate a memory-intensive object.
type Data struct {
	Value int
}

func createData() *Data {
	return &Data{
        Value: 42,
    }
}

func main() {
	for i := 0; i < 1000000; i++ {
		obj := createData()
        _ = obj
	}
	fmt.Println("Done")
}
