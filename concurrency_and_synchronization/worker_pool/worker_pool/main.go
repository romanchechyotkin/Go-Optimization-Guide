package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"runtime"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- [32]byte) {
	log.Printf("worker %d started\n", id)

	for j := range jobs {
		results <- doWork(j)
	}
}

func doWork(n int) [32]byte {
	data := []byte(fmt.Sprintf("payload-%d", n))
	return sha256.Sum256(data) //
}

// The optimal number of workers in a pool is closely tied to the number of CPU cores,
// which you can obtain in Go using runtime.NumCPU() or runtime.GOMAXPROCS(0).
// For CPU-bound tasks—where each worker consumes substantial CPU time—you generally want the number of workers to be equal to or slightly less than the number of logical CPU cores.
// This ensures maximum core utilization without excessive overhead.

// If your tasks are I/O-bound (e.g., network calls, disk I/O, database queries),
// the pool size can be larger than the number of cores.
// This is because workers will spend much of their time blocked, allowing others to run.
// In contrast, CPU-heavy workloads benefit from a smaller, tightly bounded pool that avoids contention and context switching.
func main() {
	cpu := runtime.NumCPU()
	log.Println("CPU", cpu)

	jobs := make(chan int, 100)
	results := make(chan [32]byte, 100)
	workersCount := 5

	var wg sync.WaitGroup

	for w := 1; w <= workersCount; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(w, jobs, results)
		}()
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 10; a++ {
		<-results
	}

	wg.Wait()
}
