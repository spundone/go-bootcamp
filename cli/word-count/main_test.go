package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCountFromReader(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Counts
	}{
		{
			name:     "empty",
			input:    "",
			expected: Counts{lines: 0, words: 0, chars: 0},
		},
		{
			name:     "single line",
			input:    "hello world",
			expected: Counts{lines: 1, words: 2, chars: 12}, // 11 chars + 1 newline
		},
		{
			name:     "multiple lines",
			input:    "hello\nworld\n",
			expected: Counts{lines: 2, words: 2, chars: 12},
		},
		{
			name:     "multiple words per line",
			input:    "the quick brown\nfox jumps over\n",
			expected: Counts{lines: 2, words: 6, chars: 29},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			got := countFromReader(reader, "test")
			if got.lines != tt.expected.lines {
				t.Errorf("lines = %v, want %v", got.lines, tt.expected.lines)
			}
			if got.words != tt.expected.words {
				t.Errorf("words = %v, want %v", got.words, tt.expected.words)
			}
			if got.chars != tt.expected.chars {
				t.Errorf("chars = %v, want %v", got.chars, tt.expected.chars)
			}
		})
	}
}

func TestProcessFiles(t *testing.T) {
	// Create temporary test files
	tmpDir := t.TempDir()

	normalFile := filepath.Join(tmpDir, "normal.txt")
	err := os.WriteFile(normalFile, []byte("hello\nworld\n"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	protectedFile := filepath.Join(tmpDir, "protected.txt")
	err = os.WriteFile(protectedFile, []byte("secret"), 0000)
	if err != nil {
		t.Fatal(err)
	}

	opts := Options{lines: true, words: true, chars: true}
	files := []string{
		normalFile,
		protectedFile,
		filepath.Join(tmpDir, "nonexistent.txt"),
		tmpDir, // directory
	}

	results := processFiles(files, opts)

	// Verify results
	if len(results) != len(files) {
		t.Errorf("got %d results, want %d", len(results), len(files))
	}

	// Check normal file
	if results[0].err != nil {
		t.Errorf("normal file: unexpected error: %v", results[0].err)
	}
	if results[0].lines != 2 {
		t.Errorf("normal file: got %d lines, want 2", results[0].lines)
	}

	// Check protected file
	if results[1].err == nil {
		t.Error("protected file: expected error, got nil")
	}

	// Check nonexistent file
	if results[2].err == nil {
		t.Error("nonexistent file: expected error, got nil")
	}

	// Check directory
	if results[3].err == nil {
		t.Error("directory: expected error, got nil")
	}
}

func TestParseFlags(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected Options
	}{
		{
			name:     "no flags",
			args:     []string{},
			expected: Options{lines: false, words: false, chars: false},
		},
		{
			name:     "all flags",
			args:     []string{"-l", "-w", "-c"},
			expected: Options{lines: true, words: true, chars: true},
		},
		{
			name:     "lines only",
			args:     []string{"-l"},
			expected: Options{lines: true, words: false, chars: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flags for each test
			os.Args = append([]string{"wc"}, tt.args...)
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

			got := parseFlags()
			if got.lines != tt.expected.lines ||
				got.words != tt.expected.words ||
				got.chars != tt.expected.chars {
				t.Errorf("parseFlags() = %+v, want %+v", got, tt.expected)
			}
		})
	}
}
