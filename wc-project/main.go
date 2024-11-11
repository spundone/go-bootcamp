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

// Move flag definitions to package level
var (
    flagLines = flag.Bool("l", false, "count lines")
    flagWords = flag.Bool("w", false, "count words")
    flagChars = flag.Bool("c", false, "count characters")
)

func parseFlags() Options {
    flag.Parse()
    return Options{
        lines: *flagLines,
        words: *flagWords,
        chars: *flagChars,
    }
}

	// If no files are specified, read from stdin
	if flag.NArg() == 0 {
		counts := countFromReader(os.Stdin, "")
		printCounts(counts, opts)
		if counts.err != nil {
			os.Exit(1)
		}
		return
	}

	// Process multiple files
	results := processFiles(flag.Args(), opts)
	
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

	// Print total if there are multiple files
	if len(flag.Args()) > 1 {
		total.name = "total"
		printCounts(total, opts)
	}

	if hasErrors {
		os.Exit(1)
	}
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

func countFromReader(r io.Reader, name string) Counts {
	var counts Counts
	counts.name = name

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		counts.lines++
		counts.words += countWords(line)
		counts.chars += len(line) + 1 // +1 for newline
	}

	if err := scanner.Err(); err != nil {
		counts.err = err
		return counts
	}

	return counts
}

func countWords(line string) int {
	inWord := false
	wordCount := 0

	for _, r := range line {
		if isSpace(r) {
			inWord = false
		} else if !inWord {
			wordCount++
			inWord = true
		}
	}

	return wordCount
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func printCounts(c Counts, opts Options) {
	if opts.lines {
		fmt.Printf("%8d", c.lines)
	}
	if opts.words {
		fmt.Printf("%8d", c.words)
	}
	if opts.chars {
		fmt.Printf("%8d", c.chars)
	}
	if c.name != "" {
		fmt.Printf(" %s", c.name)
	}
	fmt.Println()
} 