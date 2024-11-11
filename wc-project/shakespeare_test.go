package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestShakespeareFiles(t *testing.T) {
	// Create a temporary directory for Shakespeare files
	tmpDir, err := ioutil.TempDir("", "shakespeare-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Shakespeare test content (shortened versions)
	plays := map[string]string{
		"hamlet.txt": `To be, or not to be, that is the question:
Whether 'tis nobler in the mind to suffer
The slings and arrows of outrageous fortune,
Or to take Arms against a Sea of troubles`,

		"macbeth.txt": `Tomorrow, and tomorrow, and tomorrow,
Creeps in this petty pace from day to day,
To the last syllable of recorded time;
And all our yesterdays have lighted fools`,

		"romeo.txt": `O Romeo, Romeo, wherefore art thou Romeo?
Deny thy father and refuse thy name.
Or if thou wilt not, be but sworn my love,
And I'll no longer be a Capulet.`,
	}

	// Create test files
	for name, content := range plays {
		path := filepath.Join(tmpDir, name)
		if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Test cases for different counting options
	tests := []struct {
		name     string
		flags    []string
		expected map[string]struct {
			lines int
			words int
			chars int
		}
	}{
		{
			name:  "count all metrics",
			flags: []string{"-l", "-w", "-c"},
			expected: map[string]struct {
				lines int
				words int
				chars int
			}{
				"hamlet.txt":  {4, 40, 182},
				"macbeth.txt": {4, 33, 159},
				"romeo.txt":   {4, 37, 164},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for file, expected := range tt.expected {
				filePath := filepath.Join(tmpDir, file)
				args := append(tt.flags, filePath)
				
				// Reset flags before each test
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
				
				output := runWc(t, args...)
				expectedOutput := fmt.Sprintf("       %d       %d      %d %s\n",
					expected.lines, expected.words, expected.chars, filePath)
				
				if output != expectedOutput {
					t.Errorf("\nFile: %s\ngot:  %q\nwant: %q", file, output, expectedOutput)
				}
			}
		})
	}
}

// Optional: Add a test for actual Shakespeare files if available
func TestActualShakespeareFiles(t *testing.T) {
	shakespearePath := "shakespeare-db" // Update this path as needed
	if _, err := os.Stat(shakespearePath); os.IsNotExist(err) {
		t.Skip("Shakespeare database not available, skipping test")
	}

	files, err := ioutil.ReadDir(shakespearePath)
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".txt" {
			continue
		}

		t.Run(file.Name(), func(t *testing.T) {
			filePath := filepath.Join(shakespearePath, file.Name())
			
			// Reset flags
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			
			// Run wc with all metrics
			output := runWc(t, "-l", "-w", "-c", filePath)
			
			// We're not checking specific numbers here, just ensuring it runs
			if output == "" {
				t.Errorf("No output for file %s", file.Name())
			}
			
			// Optional: Compare with system wc command
			// This requires executing the system command and comparing results
		})
	}
} 