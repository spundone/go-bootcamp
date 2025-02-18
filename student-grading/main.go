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

	for _, record := range records {
		students = append(students, student{
			firstName: record[0],
			lastName:  record[1],
			university: record[2],
			test1Score: strconv.Atoi(record[3]),
			test2Score: strconv.Atoi(record[4]),
			test3Score: strconv.Atoi(record[5]),
			test4Score: strconv.Atoi(record[6]),
		})
	}

	return students
}

func calculateGrade(students []student) []studentStat {
	gradedStudents := make([]studentStat, 0)

	for _, student := range students {
		finalScore := float32(student.test1Score + student.test2Score + student.test3Score + student.test4Score) / 4
		grade := calculateGrade(finalScore)
		gradedStudents = append(gradedStudents, studentStat{
			student:    student,
			finalScore: finalScore,
			grade:      grade,
		})
	}

	return gradedStudents
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	return studentStat{}
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	return nil
}
