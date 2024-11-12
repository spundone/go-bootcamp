package main

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

type GrepOptions struct {
	Pattern           string
	CaseInsensitive   bool
	Recursive         bool
	LinesAfter       int
	LinesBefore      int
	CountOnly        bool
	OutputFile       string
}

type Match struct {
	LineNum     int
	Line        string
	Filename    string
	BeforeCtx   []string
	AfterCtx    []string
}

type Grep struct {
	options GrepOptions
	pattern string
}

func NewGrep(options GrepOptions) (*Grep, error) {
	pattern := options.Pattern
	if options.CaseInsensitive {
		pattern = "(?i)" + pattern
	}
	
	return &Grep{
		options: options,
		pattern: pattern,
	}, nil
}

func (g *Grep) Search(reader io.Reader, filename string) ([]Match, error) {
	scanner := bufio.NewScanner(reader)
	var matches []Match
	var lineNum int
	var contextBuffer []string
	
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		
		if matched, _ := regexp.MatchString(g.pattern, line); matched {
			match := Match{
				LineNum:  lineNum,
				Line:    line,
				Filename: filename,
			}
			
			// Handle context lines
			if g.options.LinesBefore > 0 {
				start := max(0, len(contextBuffer)-g.options.LinesBefore)
				match.BeforeCtx = append([]string{}, contextBuffer[start:]...)
			}
			
			matches = append(matches, match)
		}
		
		// Keep context buffer updated
		if g.options.LinesBefore > 0 {
			contextBuffer = append(contextBuffer, line)
			if len(contextBuffer) > g.options.LinesBefore {
				contextBuffer = contextBuffer[1:]
			}
		}
	}
	
	return matches, scanner.Err()
} 