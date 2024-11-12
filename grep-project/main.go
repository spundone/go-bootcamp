package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	opts := parseFlags()
	
	grep, err := NewGrep(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	
	// Handle standard input if no files specified
	if flag.NArg() == 0 {
		matches, err := grep.Search(os.Stdin, "")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		printMatches(matches, opts)
		return
	}
	
	// Process files
	for _, path := range flag.Args() {
		if opts.Recursive {
			err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					return processFile(grep, path, opts)
				}
				return nil
			})
		} else {
			err = processFile(grep, path, opts)
		}
		
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error processing %s: %v\n", path, err)
		}
	}
}

func processFile(grep *Grep, path string, opts GrepOptions) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	
	matches, err := grep.Search(file, path)
	if err != nil {
		return err
	}
	
	if opts.CountOnly {
		fmt.Printf("%s: %d\n", path, len(matches))
		return nil
	}
 
	printMatches(matches, opts)
	return nil
} 