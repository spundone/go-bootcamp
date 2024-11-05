package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name       string
	university string
	grades     []int
	average    float64
}

// String implements the Stringer interface
func (s Student) String() string {
	return fmt.Sprintf("%s from %s with average %.2f", s.name, s.university, s.average)
}

func readAndParseCSV(filename string) ([]Student, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var students []Student
	scanner := bufio.NewScanner(file)
	// Skip header line
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		student, err := parseStudent(line)
		if err != nil {
			return nil, fmt.Errorf("error parsing line: %v", err)
		}
		students = append(students, student)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return students, nil
}

func parseStudent(line string) (Student, error) {
	fields := strings.Split(line, ",")
	if len(fields) < 4 {
		return Student{}, fmt.Errorf("invalid number of fields")
	}

	grades := make([]int, len(fields)-2)
	var sum int
	for i, grade := range fields[2:] {
		g, err := strconv.Atoi(strings.TrimSpace(grade))
		if err != nil {
			return Student{}, fmt.Errorf("invalid grade: %v", err)
		}
		grades[i] = g
		sum += g
	}

	average := float64(sum) / float64(len(grades))
	return Student{
		name:       strings.TrimSpace(fields[0]),
		university: strings.TrimSpace(fields[1]),
		grades:     grades,
		average:    average,
	}, nil
}

func findOverallTopper(students []Student) (Student, error) {
	if len(students) == 0 {
		return Student{}, fmt.Errorf("no students found")
	}

	topper := students[0]
	for _, student := range students[1:] {
		if student.average > topper.average {
			topper = student
		}
	}
	return topper, nil
}

func findTopperPerUniversity(students []Student) (map[string]Student, error) {
	if len(students) == 0 {
		return nil, fmt.Errorf("no students found")
	}

	// Group students by university
	universitiesMap := make(map[string][]Student)
	for _, student := range students {
		universitiesMap[student.university] = append(universitiesMap[student.university], student)
	}

	// Find topper for each university
	result := make(map[string]Student)
	for university, universityStudents := range universitiesMap {
		topper, err := findOverallTopper(universityStudents)
		if err != nil {
			return nil, fmt.Errorf("error finding topper for %s: %v", university, err)
		}
		result[university] = topper
	}

	return result, nil
}

func main() {
	students, err := readAndParseCSV("students.csv")
	if err != nil {
		fmt.Printf("Error reading CSV: %v\n", err)
		return
	}

	topper, err := findOverallTopper(students)
	if err != nil {
		fmt.Printf("Error finding overall topper: %v\n", err)
		return
	}
	fmt.Printf("Overall topper: %s\n", topper)

	universityToppers, err := findTopperPerUniversity(students)
	if err != nil {
		fmt.Printf("Error finding university toppers: %v\n", err)
		return
	}
	for uni, student := range universityToppers {
		fmt.Printf("Topper for %s: %s\n", uni, student)
	}
}
