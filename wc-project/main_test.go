package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// createTestFiles creates a temporary directory with test files
// Returns the directory path and a cleanup function
func createTestFiles(t *testing.T) (string, func()) {
	tmpDir, err := ioutil.TempDir("", "wc-test")
	if err != nil {
		t.Fatal(err)
	}

	// Define test files and their content
	testFiles := map[string]string{
		"single.txt":     "Hello World\n",
		"multiple.txt":   "Line 1\nLine 2\nLine 3\n",
		"empty.txt":      "",
		"spaces.txt":     "   \n  \n    \n",
		"unicode.txt":    "Hello 世界\n你好 World\n",
		"protected.txt":  "Protected content\n",
		"whitespace.txt": "Word1  Word2   Word3\n",
	}

	// Create each test file
	for name, content := range testFiles {
		path := filepath.Join(tmpDir, name)
		if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Create a protected file for permission tests
	protectedPath := filepath.Join(tmpDir, "protected.txt")
	if err := os.Chmod(protectedPath, 0000); err != nil {
		t.Fatal(err)
	}

	// Create a directory for directory test cases
	if err := os.Mkdir(filepath.Join(tmpDir, "testdir"), 0755); err != nil {
		t.Fatal(err)
	}

	// Return cleanup function
	cleanup := func() {
		os.Chmod(protectedPath, 0644) // Reset permissions to allow deletion
		os.RemoveAll(tmpDir)
	}

	return tmpDir, cleanup
}

// TestWcSingleFile tests basic word count functionality
// Tests different flag combinations on a single file
func TestWcSingleFile(t *testing.T) {
	tmpDir, cleanup := createTestFiles(t)
	defer cleanup()

	tests := []struct {
		name     string
		flags    []string
		file     string
		expected string
	}{
		{
			name:     "count lines only",
			flags:    []string{"-l"},
			file:     "multiple.txt",
			expected: "       3 %s\n",
		},
		{
			name:     "count words only",
			flags:    []string{"-w"},
			file:     "multiple.txt",
			expected: "       6 %s\n",
		},
		{
			name:     "count chars only",
			flags:    []string{"-c"},
			file:     "multiple.txt",
			expected: "      18 %s\n",
		},
		{
			name:     "count all",
			flags:    []string{"-l", "-w", "-c"},
			file:     "multiple.txt",
			expected: "       3       6      18 %s\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flags for each test
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			
			filePath := filepath.Join(tmpDir, tt.file)
			args := append(tt.flags, filePath)
			output := runWc(t, args...)
			expected := fmt.Sprintf(tt.expected, filePath)
			
			if output != expected {
				t.Errorf("\nExpected: %q\nGot: %q", expected, output)
			}
		})
	}
}

// TestWcErrors tests error handling
// Verifies proper behavior for non-existent files, directories, and permission issues
func TestWcErrors(t *testing.T) {
	tmpDir, cleanup := createTestFiles(t)
	defer cleanup()

	tests := []struct {
		name     string
		file     string
		expected string
	}{
		{
			name:     "non-existent file",
			file:     "nonexistent.txt",
			expected: "no such file or directory",
		},
		{
			name:     "directory",
			file:     "testdir",
			expected: "is a directory",
		},
		{
			name:     "protected file",
			file:     "protected.txt",
			expected: "permission denied",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stderr output
			oldStderr := os.Stderr
			r, w, _ := os.Pipe()
			os.Stderr = w

			// Run the command
			filePath := filepath.Join(tmpDir, tt.file)
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			
			// Execute with error handling
			func() {
				defer func() {
					if r := recover(); r != nil {
						// Expected program exit
					}
				}()
				runWc(t, filePath)
			}()

			// Restore stderr and check output
			w.Close()
			os.Stderr = oldStderr
			errOutput, _ := ioutil.ReadAll(r)

			if !strings.Contains(strings.ToLower(string(errOutput)), 
				strings.ToLower(tt.expected)) {
				t.Errorf("Expected error containing %q, got %q", 
					tt.expected, string(errOutput))
			}
		})
	}
}