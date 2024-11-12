package main

import (
	"flag"
)

func parseFlags() GrepOptions {
	caseInsensitive := flag.Bool("i", false, "Case-insensitive search")
	recursive := flag.Bool("r", false, "Recursive search")
	linesAfter := flag.Int("A", 0, "Print N lines after match")
	linesBefore := flag.Int("B", 0, "Print N lines before match")
	context := flag.Int("C", 0, "Print N lines before and after match")
	countOnly := flag.Bool("c", false, "Print only count of matching lines")
	outputFile := flag.String("o", "", "Write output to file")

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	// If -C is specified, it sets both before and after context
	if *context > 0 {
		*linesAfter = *context
		*linesBefore = *context
	}

	return GrepOptions{
		Pattern:         flag.Arg(0),
		CaseInsensitive: *caseInsensitive,
		Recursive:       *recursive,
		LinesAfter:     *linesAfter,
		LinesBefore:    *linesBefore,
		CountOnly:      *countOnly,
		OutputFile:     *outputFile,
	}
} 