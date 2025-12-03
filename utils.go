package aoc25

import (
	"bufio"
	"os"
)

// MustReadLines reads the file and returns its lines as a slice.
// It loads all lines into memory; for large files, stream them instead.
// It panics if it fails to open or during reading the file.
func MustReadLines(name string) []string {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return lines
}
