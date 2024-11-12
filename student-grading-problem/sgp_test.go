package main

import (
	"reflect" // Package for deep comparison of values
	"testing" // Package for writing tests
)

// TestParseStudent tests the parseStudent function
func TestParseStudent(t *testing.T) {
	// Define test cases as a slice of structs
	tests := []struct {
		name    string    // Name of the test case
		line    string    // Input line to parse
		want    Student   // Expected output
		wantErr bool      // Indicates if an error is expected
	}{
		{
			name: "valid student", // Test case name
			line: "John Doe,University A,90,85,88", // Input line
			want: Student{ // Expected output
				name:       "John Doe",
				university: "University A",
				grades:     []int{90, 85, 88},
					average:    87.66666666666667,
			},
			wantErr: false, // No error expected
		},
		{
			name:    "invalid format", // Test case name
			line:    "John Doe,University A", // Input line with missing grades
			wantErr: true, // Error expected
		},
		{
			name:    "invalid grade", // Test case name
			line:    "John Doe,University A,90,invalid,88", // Input line with an invalid grade
			wantErr: true, // Error expected
		},
	}

	// Loop through each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // Run the test case
			got, err := parseStudent(tt.line) // Call the function to test
			if (err != nil) != tt.wantErr { // Check if the error matches expectation
				t.Errorf("parseStudent() error = %v, wantErr %v", err, tt.wantErr) // Report error
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) { // Check if the output matches expectation
				t.Errorf("parseStudent() = %v, want %v", got, tt.want) // Report mismatch
			}
		})
	}
}

// TestFindOverallTopper tests the findOverallTopper function
func TestFindOverallTopper(t *testing.T) {
	// Define test cases as a slice of structs
	tests := []struct {
		name    string    // Name of the test case
		students []Student // Input list of students
		want    Student   // Expected output
		wantErr bool      // Indicates if an error is expected
	}{
		{
			name: "normal case", // Test case name
			students: []Student{ // Input list of students
				{name: "John", average: 85},
				{name: "Jane", average: 90},
				{name: "Bob", average: 88},
			},
			want:    Student{name: "Jane", average: 90}, // Expected output
			wantErr: false, // No error expected
		},
		{
			name:    "empty list", // Test case name
			students: []Student{}, // Input list is empty
			want:    Student{}, // Expected output is an empty Student
			wantErr: true, // Error expected
		},
		{
			name: "single student", // Test case name
			students: []Student{ // Input list with one student
				{name: "John", average: 85},
			},
			want:    Student{name: "John", average: 85}, // Expected output
			wantErr: false, // No error expected
		},
	}

	// Loop through each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // Run the test case
			got, err := findOverallTopper(tt.students) // Call the function to test
			if (err != nil) != tt.wantErr { // Check if the error matches expectation
				t.Errorf("findOverallTopper() error = %v, wantErr %v", err, tt.wantErr) // Report error
				return
			}
			if !tt.wantErr && got.name != tt.want.name { // Check if the output matches expectation
				t.Errorf("findOverallTopper() = %v, want %v", got, tt.want) // Report mismatch
			}
		})
	}
}

// TestFindTopperPerUniversity tests the findTopperPerUniversity function
func TestFindTopperPerUniversity(t *testing.T) {
	// Define test cases as a slice of structs
	tests := []struct {
		name    string    // Name of the test case
		students []Student // Input list of students
		want    map[string]Student // Expected output map of toppers
		wantErr bool      // Indicates if an error is expected
	}{
		{
			name: "normal case", // Test case name
			students: []Student{ // Input list of students
				{name: "John", university: "Uni A", average: 85},
				{name: "Jane", university: "Uni A", average: 90},
				{name: "Bob", university: "Uni B", average: 88},
			},
			want: map[string]Student{ // Expected output
				"Uni A": {name: "Jane", university: "Uni A", average: 90},
				"Uni B": {name: "Bob", university: "Uni B", average: 88},
			},
			wantErr: false, // No error expected
		},
		{
			name:    "empty list", // Test case name
			students: []Student{}, // Input list is empty
			want:    nil, // Expected output is nil
			wantErr: true, // Error expected
		},
		{
			name: "single university", // Test case name
			students: []Student{ // Input list with students from one university
				{name: "John", university: "Uni A", average: 85},
				{name: "Jane", university: "Uni A", average: 90},
			},
			want: map[string]Student{ // Expected output
				"Uni A": {name: "Jane", university: "Uni A", average: 90},
			},
			wantErr: false, // No error expected
		},
	}

	// Loop through each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // Run the test case
			got, err := findTopperPerUniversity(tt.students) // Call the function to test
			if (err != nil) != tt.wantErr { // Check if the error matches expectation
				t.Errorf("findTopperPerUniversity() error = %v, wantErr %v", err, tt.wantErr) // Report error
				return
			}
			if !tt.wantErr { // If no error is expected
				for uni, student := range tt.want { // Loop through expected toppers
					if got[uni].name != student.name || got[uni].average != student.average { // Check if the output matches expectation
						t.Errorf("findTopperPerUniversity() for %s = %v, want %v", uni, got[uni], student) // Report mismatch
					}
				}
			}
		})
	}
} 