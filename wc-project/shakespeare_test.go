package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestShakespeareFiles(t *testing.T) {
	shakespearePath := "shakespeare-db"

	if _, err := os.Stat(shakespearePath); os.IsNotExist(err) {
		t.Skip("shakespeare-db directory not found - skipping Shakespeare tests")
	}

	files, err := ioutil.ReadDir(shakespearePath)
	if err != nil {
		t.Fatalf("Failed to read shakespeare-db directory: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".txt" {
			continue
		}

		t.Run(file.Name(), func(t *testing.T) {
			filePath := filepath.Join(shakespearePath, file.Name())
			
			sysOutput, err := getSystemWcOutput(filePath)
			if err != nil {
				t.Fatalf("Failed to get system wc output: %v", err)
			}

			ourOutput := runWc(t, "-l", "-w", "-c", filePath)

			sysNums := extractNumbers(sysOutput)
			ourNums := extractNumbers(ourOutput)

			if len(sysNums) != len(ourNums) {
				t.Errorf("Output mismatch for %s:\nSystem wc: %v\nOur wc: %v", 
					file.Name(), sysNums, ourNums)
			}

			t.Logf("File: %s\nLines: %d\nWords: %d\nChars: %d\n",
				file.Name(), ourNums[0], ourNums[1], ourNums[2])
		})
	}
}

func getSystemWcOutput(filepath string) (string, error) {
	cmd := exec.Command("wc", filepath)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

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