package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
)

type Counts struct {
	lines   int
	words   int
	chars   int
	name    string
	err     error
}

type Options struct {
	lines bool
	words bool
	chars bool
}

func main() {
	// Parse command line flags
	opts := parseFlags()

	// If no specific options are given, show all counts
	if !opts.lines && !opts.words && !opts.chars {
		opts.lines = true
		opts.words = true
		opts.chars = true
	}

	// If no files are specified, read from stdin
	if flag.NArg() == 0 {
		counts := countFromReader(os.Stdin, "")
		printCounts(counts, opts)
	}
}

// Print results and calculate total
total := Counts{}
hasErrors := false

for _, result := range results {
	if result.err != nil {
		fmt.Fprintf(os.Stderr, "./wc: %s: %v\n", result.name, result.err)
		hasErrors = true
		continue
	}
	printCounts(result, opts)
	total.lines += result.lines
	total.words += result.words
	total.chars += result.chars
}
