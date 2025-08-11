package main

import (
	"bufio"
	"os"
)

func main() {
	// without buffering
	f, _ := os.Create("output.txt")
	for i := 0; i < 10000; i++ {
		f.Write([]byte("line\n"))
	}

	// with buffering
	f, _ = os.Create("output2.txt")
	buf := bufio.NewWriter(f)
	for i := 0; i < 10000; i++ {
		buf.WriteString("line\n")
	}
	buf.Flush() // ensure all buffered data is written

}
