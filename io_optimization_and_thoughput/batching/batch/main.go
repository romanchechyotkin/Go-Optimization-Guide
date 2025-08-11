package main

import "strings"

// NOT WELL
func logLine(line string) {
	f.WriteString(line + "\n")
}

var batch []string

// WELL
func logBatch(line string) {
	batch = append(batch, line)

	if len(batch) >= 100 {
		f.WriteString(strings.Join(batch, "\n") + "\n")
		batch = batch[:0]
	}
}
