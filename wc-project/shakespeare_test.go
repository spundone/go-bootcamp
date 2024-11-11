package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// TestShakespeareFiles tests word count functionality on Shakespeare's plays
// It reads files from the shakespeare-db directory and verifies the word count
func TestShakespeareFiles(t *testing.T) {
	// Define path to Shakespeare files
	shakespearePath := "shakespeare-db"

	// Skip test if Shakespeare directory doesn't exist
	if _, err := os.Stat(shakespearePath); os.IsNotExist(err) {
		t.Skip("shakespeare-db directory not found - skipping Shakespeare tests")
	}

	// Read all files from the Shakespeare directory
	files, err := ioutil.ReadDir(shakespearePath)
	if err != nil {
		t.Fatalf("Failed to read shakespeare-db directory: %v", err)
	}

	// Process each .txt file
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".txt" {
			continue
		}

		t.Run(file.Name(), func(t *testing.T) {
			filePath := filepath.Join(shakespearePath, file.Name())
			
			// Get system wc command output for comparison
			sysOutput, err := getSystemWcOutput(filePath)
			if err != nil {
				t.Fatalf("Failed to get system wc output: %v", err)
			}

			// Get our wc implementation output
			ourOutput := runWcCommand(t, "-l", "-w", "-c", filePath)

			// Compare results (ignoring exact formatting)
			sysNums := extractNumbers(sysOutput)
			ourNums := extractNumbers(ourOutput)

			if len(sysNums) != len(ourNums) {
				t.Errorf("Output mismatch for %s:\nSystem wc: %v\nOur wc: %v", 
					file.Name(), sysNums, ourNums)
			}

			// Print the word count results
			t.Logf("File: %s\nLines: %d\nWords: %d\nChars: %d\n",
				file.Name(), ourNums[0], ourNums[1], ourNums[2])
		})
	}
}

// Helper function to run system wc command
func getSystemWcOutput(filepath string) (string, error) {
	cmd := exec.Command("wc", filepath)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// Helper function to extract numbers from wc output
func extractNumbers(output string) []int {
	var numbers []int
	fields := strings.Fields(output)
	for _, field := range fields {
		if n, err := strconv.Atoi(field); err == nil {
			numbers = append(numbers, n)
		}
	}
	return numbers
} 