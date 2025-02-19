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
	totalScore float32
	grade      Grade
}

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


func calculateGrade(totalScore float32) Grade {
	if totalScore >= 90 {
		return A
	} else if totalScore >= 80 {
		return B
	} else if totalScore >= 70 {
		return C
	} else {
		return F
	}
}

func calculateGrades(students []student) []studentStat {
	gradedStudents := make([]studentStat, 0)

	for _, student := range students {
		totalScore := float32(student.test1Score+student.test2Score+student.test3Score+student.test4Score) / 4
		grade := calculateGrade(totalScore)
		gradedStudents = append(gradedStudents, studentStat{
			student:    student,
			totalScore: totalScore,
			grade:      grade,
		})
	}

	return gradedStudents
}

func findOverallTopper(students []studentStat) studentStat {
	topper := students[0]

	for _, student := range students {
		if student.totalScore > topper.totalScore {
			topper = student
		}
	}
	return topper
}

func findTopperPerUniversity(students []studentStat) map[string]studentStat {
	toppers := make(map[string]studentStat)

	for _, student := range students {
		if _, ok := toppers[student.university]; !ok || student.totalScore > toppers[student.university].totalScore {
			toppers[student.university] = student
		}
	}
	return toppers
}

func main() {
	students := parseCSV("grades.csv")
	gradedStudents := calculateGrades(students)
	overallTopper := findOverallTopper(gradedStudents)
	toppersPerUniversity := findTopperPerUniversity(gradedStudents)

	fmt.Println("Overall Topper: ", overallTopper)
	fmt.Println("Toppers per University: ", toppersPerUniversity)
}