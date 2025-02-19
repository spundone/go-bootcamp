package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

// Sol starts here
func parseCSV(filePath string) []student {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1

	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	students := make([]student, 0)

	// Skip the header row
	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}
		test1Score, _ := strconv.Atoi(record[3]) // Convert to int
		test2Score, _ := strconv.Atoi(record[4]) // Convert to int
		test3Score, _ := strconv.Atoi(record[5]) // Convert to int
		test4Score, _ := strconv.Atoi(record[6]) // Convert to int

		students = append(students, student{
			firstName:  record[0],
			lastName:   record[1],
			university: record[2],
			test1Score: test1Score,
			test2Score: test2Score,
			test3Score: test3Score,
			test4Score: test4Score,
		})
	}

	return students
}

func calculateGrade(students []student) []studentStat {
	gradedStudents := make([]studentStat, 0)
	grade := F //by default it's set to F
	for _, student := range students {
		totalScore := float32(student.test1Score+student.test2Score+student.test3Score+student.test4Score) / 4
		if totalScore >= 70 {
			grade = A
		} else if totalScore >= 50 { //no need to define <=70
			grade = B
		} else if totalScore >= 35 {
			grade = C
		} else {
			grade = F
		}

		gradedStudents = append(gradedStudents, studentStat{
			student:    student,
			finalScore: totalScore,
			grade:      grade,
		})
	}
	return gradedStudents
}

func findOverallTopper(students []studentStat) studentStat {
	topper := students[0]

	for _, student := range students {
		if student.finalScore > topper.finalScore {
			topper = student
		}
	}
	return topper
}

func findTopperPerUniversity(students []studentStat) map[string]studentStat {
	toppers := make(map[string]studentStat)

	for _, student := range students {
		if _, ok := toppers[student.university]; !ok || student.finalScore > toppers[student.university].finalScore {
			toppers[student.university] = student
		}
	}
	return toppers
}

func main() {
	students := parseCSV("grades.csv")
	gradedStudents := calculateGrade(students)
	overallTopper := findOverallTopper(gradedStudents)
	toppersPerUniversity := findTopperPerUniversity(gradedStudents)
	fmt.Println("Overall Topper: ", overallTopper)
	fmt.Println("Toppers per University: ", toppersPerUniversity)
}
