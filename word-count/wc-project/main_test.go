package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Helper function to create test files
func createTestFiles(t *testing.T) (string, func()) {
	// Create temporary directory
	tmpDir, err := ioutil.TempDir("", "wc-test")
	if err != nil {
		t.Fatal(err)
	}

	// Create test files
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

	// Create protected file
	protectedPath := filepath.Join(tmpDir, "protected.txt")
	if err := os.Chmod(protectedPath, 0000); err != nil {
		t.Fatal(err)
	}

	// Create directory
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
			expected: "       3 multiple.txt\n",
		},
		{
			name:     "count words only",
			flags:    []string{"-w"},
			file:     "multiple.txt",
			expected: "       6 multiple.txt\n",
		},
		{
			name:     "count chars only",
			flags:    []string{"-c"},
			file:     "multiple.txt",
			expected: "      18 multiple.txt\n",
		},
		{
			name:     "count all",
			flags:    []string{"-l", "-w", "-c"},
			file:     "multiple.txt",
			expected: "       3       6      18 multiple.txt\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := append(tt.flags, filepath.Join(tmpDir, tt.file))
			output := runWc(t, args...)
			if output != tt.expected {
				t.Errorf("got %q, want %q", output, tt.expected)
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
			expected: "./wc: nonexistent.txt: open: no such file or directory\n",
		},
		{
			name:     "directory",
			file:     "testdir",
			expected: "./wc: testdir: read: is a directory\n",
		},
		{
			name:     "protected file",
			file:     "protected.txt",
			expected: "./wc: protected.txt: open: permission denied\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := runWc(t, filepath.Join(tmpDir, tt.file))
			if !strings.Contains(strings.ToLower(output), strings.ToLower(tt.expected)) {
				t.Errorf("got %q, want %q", output, tt.expected)
			}
		})
	}
}

func TestWcMultipleFiles(t *testing.T) {
	tmpDir, cleanup := createTestFiles(t)
	defer cleanup()

	files := []string{
		filepath.Join(tmpDir, "single.txt"),
		filepath.Join(tmpDir, "multiple.txt"),
	}

	output := runWc(t, files...)
	expected := `       1       2      12 single.txt
       3       6      18 multiple.txt
       4       8      30 total
`

	if output != expected {
		t.Errorf("got %q, want %q", output, expected)
	}
}

// Helper function to run wc command
func runWc(t *testing.T, args ...string) string {
	t.Helper()
	
	// Save original args and restore them after the test
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Set up command line arguments
	os.Args = append([]string{"wc"}, args...)

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run main
	main()

	// Restore stdout and read output
	w.Close()
	os.Stdout = oldStdout
	output, _ := ioutil.ReadAll(r)

	return string(output)
} 