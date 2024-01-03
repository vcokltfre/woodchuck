package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <start> <end>\n", os.Args[0])
		os.Exit(1)
	}

	start := os.Args[2]
	end := os.Args[3]

	full := start + end

	qualifier := "a"
	if start[0] == 'a' || start[0] == 'e' || start[0] == 'i' || start[0] == 'o' || start[0] == 'u' {
		qualifier = "an"
	}

	fmt.Printf("How much %s could %s %s %s if %s %s could %s %s?\n", start, qualifier, full, end, qualifier, full, end, start)
}
