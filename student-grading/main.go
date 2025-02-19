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
	graded := F //by default it's set to F
	for _, s := range students {
		totalScore := float32(s.test1Score+s.test2Score+s.test3Score+s.test4Score) / 4
		if totalScore >= 70 {
			graded = A
		} else if totalScore >= 50 { //no need to define <=70 since its handled above
			graded = B
		} else if totalScore >= 35 {
			graded = C
		} else {
			graded = F
		}

		gradedStudents = append(gradedStudents, studentStat{
			student:    s,
			finalScore: totalScore,
			grade:      graded,
		})
	}
	return gradedStudents
}

func findOverallTopper(students []studentStat) studentStat { 
	topperO := students[0]

	for _, s := range students {
		if s.finalScore > topperO.finalScore {
			topperO = s
		}
	}
	return topperO
}

func findTopperPerUniversity(students []studentStat) map[string]studentStat {
	topperU := make(map[string]studentStat)

	for _, s := range students {
		if _, ok := topperU[s.university]; !ok || s.finalScore > topperU[s.university].finalScore {
			topperU[s.university] = s
		}
	}
	return topperU
}

func main() {
	students := parseCSV("grades.csv")
	gradedStudents := calculateGrade(students)
	overallTopper := findOverallTopper(gradedStudents)
	toppersPerUniversity := findTopperPerUniversity(gradedStudents)
	fmt.Println("Overall Topper: ", overallTopper)
	fmt.Println("Toppers per University: ", toppersPerUniversity)
}
