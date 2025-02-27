package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

// Options holds the command line flags
type Options struct {
	lines bool
	words bool
	chars bool
}

// Counts holds the counting results for a file
type Counts struct {
	name  string
	lines int
	words int
	chars int
	err   error
}

// countFromReader counts lines, words and characters from a reader
func countFromReader(r io.Reader, name string) Counts {
	scanner := bufio.NewScanner(r)
	counts := Counts{name: name}

	for scanner.Scan() {
		line := scanner.Text()
		counts.lines++
		counts.words += len(strings.Fields(line))
		counts.chars += len(line) + 1 // +1 for the newline character
	}

	if err := scanner.Err(); err != nil {
		counts.err = fmt.Errorf("read: %v", err)
	}
	return counts
}

func parseFlags() Options {
	lines := flag.Bool("l", false, "count lines")
	words := flag.Bool("w", false, "count words")
	chars := flag.Bool("c", false, "count characters")
	flag.Parse()

	return Options{
		lines: *lines,
		words: *words,
		chars: *chars,
	}
}

func processFiles(filenames []string, opts Options) []Counts {
	results := make([]Counts, len(filenames))
	var wg sync.WaitGroup

	// Process each file in a separate goroutine
	for i, filename := range filenames {
		wg.Add(1)
		go func(idx int, fname string) {
			defer wg.Done()

			file, err := os.Open(fname)
			if err != nil {
				results[idx] = Counts{name: fname, err: err}
				return
			}
			defer file.Close()

			results[idx] = countFromReader(file, fname)
		}(i, filename)
	}

	wg.Wait()
	return results
}

func main() {
	opts := parseFlags()
	args := flag.Args()

	// If no options specified, show all counts
	if !opts.lines && !opts.words && !opts.chars {
		opts.lines = true
		opts.words = true
		opts.chars = true
	}

	// If no files specified, read from stdin
	if len(args) == 0 {
		counts := countFromReader(os.Stdin, "")
		if counts.err != nil {
			fmt.Fprintf(os.Stderr, "./wc: %v\n", counts.err)
			os.Exit(1)
		}
		printCounts(counts, opts)
		return
	}

	// Process all input files
	results := processFiles(args, opts)
	total := Counts{}
	hasErrors := false

	// Print results and collect totals
	for _, counts := range results {
		if counts.err != nil {
			fmt.Fprintf(os.Stderr, "./wc: %s: %v\n", counts.name, counts.err)
			hasErrors = true
			continue
		}
		printCounts(counts, opts)
		total.lines += counts.lines
		total.words += counts.words
		total.chars += counts.chars
	}

	// Print totals if more than one file
	if len(args) > 1 {
		total.name = "total"
		printCounts(total, opts)
	}

	if hasErrors {
		os.Exit(1)
	}
}

func printCounts(counts Counts, opts Options) {
	format := ""
	values := []interface{}{}

	if opts.lines {
		format += "%8d"
		values = append(values, counts.lines)
	}
	if opts.words {
		format += "%8d"
		values = append(values, counts.words)
	}
	if opts.chars {
		format += "%8d"
		values = append(values, counts.chars)
	}

	if counts.name != "" {
		format += " %s"
		values = append(values, counts.name)
	}

	fmt.Printf(format+"\n", values...)
}
