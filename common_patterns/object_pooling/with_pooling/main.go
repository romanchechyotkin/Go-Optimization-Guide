package main

import (
    "fmt"
    "sync"
)

type Data struct {
    Value int
}

var dataPool = sync.Pool{
    New: func() any {
        return &Data{}
    },
}

func main() {
    for i := 0; i < 1000000; i++ {
        obj := dataPool.Get().(*Data) // Retrieve from pool
        obj.Value = 42 // Use the object
        dataPool.Put(obj) // Return object to pool for reuse
    }
    fmt.Println("Done")
}
