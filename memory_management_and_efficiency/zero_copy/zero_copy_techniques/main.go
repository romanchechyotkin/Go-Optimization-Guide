package main

import (
	"io"
	"log"
)

func StreamData(src io.Reader, dst io.Writer) error {
	buf := make([]byte, 4096) // Reusable buffer
	_, err := io.CopyBuffer(dst, src, buf)
	return err
}

func process(buffer []byte) []byte {
	return buffer[128:256] // returns a slice reference without copying
}

func main() {
	l := 256
	b := make([]byte, l)
	for i := range l {
		b[i] = byte(i)
	}
	log.Printf("origin b=%v; len=%d; cap=%d\n", b, len(b), cap(b))

	b2 := process(b)
	log.Printf("new b2 without copying original b=%v; len=%d; cap=%d\n", b2, len(b2), cap(b2))
	b2[len(b2)-1] = 69

	log.Printf("lat element of b=%d\n", b[len(b)-1])
	log.Printf("lat element of b2=%d\n", b2[len(b2)-1])
}
