package common

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(fileName string) []string {
	// Open file
	file, err := os.Open(fileName)
	defer file.Close()

	// Create scanner
	if err != nil { log.Fatalf("failed to open") }
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}
