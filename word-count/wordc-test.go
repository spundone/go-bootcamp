package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"empty string", "", 0},
		{"single word", "hello", 1},
		{"multiple words", "hello world", 2},
		{"multiple spaces", "hello   world", 2},
		{"tabs", "hello\tworld", 2},
		{"mixed whitespace", "hello \t world", 2},
		{"emoji", "hello 👋 world", 3},
		{"only whitespace", "   \t   ", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countWords(tt.input)
			if got != tt.expected {
				t.Errorf("countWords(%q) = %d; want %d", tt.input, got, tt.expected)
			}
		})
	}
}

func TestCountFromReader(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedLines int
		expectedWords int
		expectedChars int
	}{
		{
			name:          "empty file",
			input:         "",
			expectedLines: 0,
			expectedWords: 0,
			expectedChars: 0,
		},
		{
			name:          "single line",
			input:         "hello world\n",
			expectedLines: 1,
			expectedWords: 2,
			expectedChars: 12,
		},
		{
			name:          "multiple lines",
			input:         "hello world\ntest line\n",
			expectedLines: 2,
			expectedWords: 4,
			expectedChars: 21,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			counts := countFromReader(reader, "")

			if counts.err != nil {
				t.Errorf("unexpected error: %v", counts.err)
			}
			if counts.lines != tt.expectedLines {
				t.Errorf("lines = %d; want %d", counts.lines, tt.expectedLines)
			}
			if counts.words != tt.expectedWords {
				t.Errorf("words = %d; want %d", counts.words, tt.expectedWords)
			}
			if counts.chars != tt.expectedChars {
				t.Errorf("chars = %d; want %d", counts.chars, tt.expectedChars)
			}
		})
	}
}

func TestProcessFiles(t *testing.T) {
	// Create temporary directory
	tmpDir, err := ioutil.TempDir("", "wc-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test files
	files := map[string]string{
		"file1.txt": "hello world\ntest line\n",
		"file2.txt": "another\ntest\nfile\n",
	}

	for name, content := range files {
		path := filepath.Join(tmpDir, name)
		if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Create directory for error testing
	dirPath := filepath.Join(tmpDir, "testdir")
	if err := os.Mkdir(dirPath, 0755); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name      string
		files     []string
		wantError bool
	}{
		{
			name:      "valid files",
			files:     []string{filepath.Join(tmpDir, "file1.txt"), filepath.Join(tmpDir, "file2.txt")},
			wantError: false,
		},
		{
			name:      "non-existent file",
			files:     []string{filepath.Join(tmpDir, "nonexistent.txt")},
			wantError: true,
		},
		{
			name:      "directory",
			files:     []string{dirPath},
			wantError: true,
		},
	}

	opts := Options{lines: true, words: true, chars: true}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := processFiles(tt.files, opts)
			
			hasError := false
			for _, result := range results {
				if result.err != nil {
					hasError = true
					break
				}
			}

			if hasError != tt.wantError {
				t.Errorf("processFiles() error = %v, wantError %v", hasError, tt.wantError)
			}
		})
	}
} 