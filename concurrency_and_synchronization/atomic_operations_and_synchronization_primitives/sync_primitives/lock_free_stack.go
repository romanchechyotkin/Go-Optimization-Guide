package main

import "sync/atomic"

// lock free stack

type node struct {
	next *node
	val  any
}

var head atomic.Pointer[node]

// 1 => [] => [1]
// 2 => [1] => [1,2]

func push(n *node) {
	for {
		old := head.Load()
		n.next = old
		if head.CompareAndSwap(old, n) {
			return
		}
	}
}
