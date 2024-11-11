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

// runWc executes the wc command with given arguments and returns the output
func runWc(t *testing.T, args ...string) string {
	t.Helper()
	
	// Save original args and flags
	oldArgs := os.Args
	oldFlags := flag.CommandLine
	defer func() { 
		os.Args = oldArgs 
		flag.CommandLine = oldFlags
	}()

	// Set up new args
	os.Args = append([]string{"wc"}, args...)
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	// Run main with panic recovery
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Recovered from panic: %v", r)
		}
	}()

	main()

	// Get output
	w.Close()
	output, _ := ioutil.ReadAll(r)
	return string(output)
}

// Helper function to create test files
func createTestFiles(t *testing.T) (string, func()) {
	tmpDir, err := ioutil.TempDir("", "wc-test")
	if err != nil {
		t.Fatal(err)
	}

	testFiles := map[string]string{
		"single.txt":     "Hello World\n",
		"multiple.txt":   "Line 1\nLine 2\nLine 3\n",
		"empty.txt":      "",
		"spaces.txt":     "   \n  \n    \n",
		"unicode.txt":    "Hello 世界\n你好 World\n",
		"protected.txt":  "Protected content\n",
		"whitespace.txt": "Word1  Word2   Word3\n",
	}

	for name, content := range testFiles {
		path := filepath.Join(tmpDir, name)
		if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	protectedPath := filepath.Join(tmpDir, "protected.txt")
	if err := os.Chmod(protectedPath, 0000); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(tmpDir, "testdir"), 0755); err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		os.Chmod(protectedPath, 0644)
		os.RemoveAll(tmpDir)
	}

	return tmpDir, cleanup
}

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
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
		// Add more error test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStderr := os.Stderr
			r, w, _ := os.Pipe()
			os.Stderr = w
			defer func() { os.Stderr = oldStderr }()

			filePath := filepath.Join(tmpDir, tt.file)
			_ = runWc(t, filePath)

			w.Close()
			errOutput, _ := ioutil.ReadAll(r)

			if !strings.Contains(strings.ToLower(string(errOutput)), strings.ToLower(tt.expected)) {
				t.Errorf("Expected error containing %q, got %q", tt.expected, string(errOutput))
			}
		})
	}
}