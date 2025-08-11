package main

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

var requests atomic.Int64

func handleRequest() {
	requests.Add(1)
}

var shutdown atomic.Int32

func mainLoop() {
	for {
		if shutdown.Load() == 1 {
			break
		}
		// do work
	}
}

func stop() {
	shutdown.Store(1)
}

var resource unsafe.Pointer
var initStatus int32 // 0: not started, 1: in progress, 2: completed

func getResource() *MyResource {
	if atomic.LoadInt32(&initStatus) == 2 {
		return (*MyResource)(atomic.LoadPointer(&resource))
	}

	if atomic.CompareAndSwapInt32(&initStatus, 0, 1) {
		newRes := expensiveInit() // initialization logic
		atomic.StorePointer(&resource, unsafe.Pointer(newRes))
		atomic.StoreInt32(&initStatus, 2)
		return newRes
	}

	for atomic.LoadInt32(&initStatus) != 2 {
		runtime.Gosched() // yield until the initializer finishes
	}

	return (*MyResource)(atomic.LoadPointer(&resource))
}
