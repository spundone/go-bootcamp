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
			// Reset flags before each test
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			
			filePath := filepath.Join(tmpDir, tt.file)
			args := append(tt.flags, filePath)
			output := runWc(t, args...)
			expected := fmt.Sprintf(tt.expected, filePath)
			
			if output != expected {
				t.Errorf("got %q, want %q", output, expected)
			}
		})
	}
}

// Rest of the test file remains the same...

func runWc(t *testing.T, args ...string) string {
	t.Helper()
	
	// Reset flags and args
	oldArgs := os.Args
	oldFlags := flag.CommandLine
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	
	defer func() { 
		os.Args = oldArgs 
		flag.CommandLine = oldFlags
	}()

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