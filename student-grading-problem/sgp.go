package main

import (
	"bufio" // Package for buffered I/O
	"fmt"   // Package for formatted I/O
	"os"    // Package for OS functions (like file handling)
	"strconv" // Package for string conversions
	"strings" // Package for string manipulation
)

// Student struct holds information about a student
type Student struct {
	name       string   // Name of the student
	university string   // University the student attends
	grades     []int    // Slice to hold grades
	average    float64  // Average grade of the student
}

// String implements the Stringer interface for pretty printing
func (s Student) String() string {
	return fmt.Sprintf("%s from %s with average %.2f", s.name, s.university, s.average)
}

// readAndParseCSV reads student data from a CSV file and returns a slice of Student
func readAndParseCSV(filename string) ([]Student, error) {
	file, err := os.Open(filename) // Open the CSV file
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err) // Return error if file can't be opened
	}
	defer file.Close() // Ensure the file is closed when the function exits

	var students []Student // Slice to hold the students
	scanner := bufio.NewScanner(file) // Create a new scanner to read the file
	// Skip header line
	scanner.Scan() // Read the first line (header) and ignore it

	// Read each line of the file
	for scanner.Scan() {
		line := scanner.Text() // Get the current line
		student, err := parseStudent(line) // Parse the line into a Student struct
		if err != nil {
			return nil, fmt.Errorf("error parsing line: %v", err) // Return error if parsing fails
		}
		students = append(students, student) // Add the student to the slice
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err) // Return error if reading fails
	}

	return students, nil // Return the slice of students
}

// parseStudent takes a CSV line and converts it into a Student struct
func parseStudent(line string) (Student, error) {
	fields := strings.Split(line, ",") // Split the line by commas
	if len(fields) < 4 {
		return Student{}, fmt.Errorf("invalid number of fields") // Return error if not enough fields
	}

	grades := make([]int, len(fields)-2) // Create a slice for grades
	var sum int // Variable to hold the sum of grades
	for i, grade := range fields[2:] { // Iterate over the grades in the fields
		g, err := strconv.Atoi(strings.TrimSpace(grade)) // Convert grade to integer
		if err != nil {
			return Student{}, fmt.Errorf("invalid grade: %v", err) // Return error if conversion fails
		}
		grades[i] = g // Store the grade
		sum += g // Add to the sum
	}

	// Calculate the average grade
	average := float64(sum) / float64(len(grades))
	return Student{
		name:       strings.TrimSpace(fields[0]), // Trim whitespace from name
		university: strings.TrimSpace(fields[1]), // Trim whitespace from university name
		grades:     grades, // Assign grades
		average:    average, // Assign average
	}, nil
}

// findOverallTopper finds the student with the highest average grade
func findOverallTopper(students []Student) (Student, error) {
	if len(students) == 0 {
		return Student{}, fmt.Errorf("no students found") // Return error if no students are present
	}

	topper := students[0] // Assume the first student is the topper
	for _, student := range students[1:] { // Iterate over the rest of the students
		if student.average > topper.average { // Check if the current student has a higher average
			topper = student // Update topper
		}
	}
	return topper, nil // Return the topper
}

// findTopperPerUniversity finds the topper for each university
func findTopperPerUniversity(students []Student) (map[string]Student, error) {
	if len(students) == 0 {
		return nil, fmt.Errorf("no students found") // Return error if no students are present
	}

	// Group students by university
	universitiesMap := make(map[string][]Student) // Map to hold students by university
	for _, student := range students {
		universitiesMap[student.university] = append(universitiesMap[student.university], student) // Append student to their university
	}

	// Find topper for each university
	result := make(map[string]Student) // Map to hold the toppers
	for university, universityStudents := range universitiesMap {
		topper, err := findOverallTopper(universityStudents) // Find the topper for the university
		if err != nil {
			return nil, fmt.Errorf("error finding topper for %s: %v", university, err) // Return error if finding topper fails
		}
		result[university] = topper // Store the topper in the result map
	}

	return result, nil // Return the map of university toppers
}

// main function is the entry point of the program
func main() {
	students, err := readAndParseCSV("students.csv") // Read and parse the CSV file
	if err != nil {
		fmt.Printf("Error reading CSV: %v\n", err) // Print error if reading fails
		return
	}

	topper, err := findOverallTopper(students) // Find the overall topper
	if err != nil {
		fmt.Printf("Error finding overall topper: %v\n", err) // Print error if finding topper fails
		return
	}
	fmt.Printf("Overall topper: %s\n", topper) // Print the overall topper

	universityToppers, err := findTopperPerUniversity(students) // Find toppers for each university
	if err != nil {
		fmt.Printf("Error finding university toppers: %v\n", err) // Print error if finding university toppers fails
		return
	}
	for uni, student := range universityToppers {
		fmt.Printf("Topper for %s: %s\n", uni, student) // Print each university's topper
	}
}
