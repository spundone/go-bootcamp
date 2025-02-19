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

// Implementing the Stringer interface for studentStat
func (s studentStat) String() string {
	return fmt.Sprintf("%s %s from %s: Final Score: %.2f, Grade: %s",
		s.firstName, s.lastName, s.university, s.finalScore, s.grade)
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
		var s student
		s.firstName = record[0]
		s.lastName = record[1]
		s.university = record[2]

		// Convert scores and handle errors
		for j := 3; j <= 6; j++ {
			score, err := strconv.Atoi(record[j])
			if err != nil {
				fmt.Printf("error converting score for %s %s: ", s.firstName, s.lastName)
				continue
			}
			switch j {
			case 3:
				s.test1Score = score
			case 4:
				s.test2Score = score
			case 5:
				s.test3Score = score
			case 6:
				s.test4Score = score
			}
		}
		students = append(students, s)
	}

	return students
}

func calculateGrade(students []student) []studentStat {
	gradedStudents := make([]studentStat, 0)
	for _, s := range students {
		totalScore := float32(s.test1Score+s.test2Score+s.test3Score+s.test4Score) / 4
		var graded Grade

		switch {
		case totalScore >= 70:
			graded = A
		case totalScore >= 50:
			graded = B
		case totalScore >= 35:
			graded = C
		default:
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
	if len(students) == 0 {
		return studentStat{} // Return zero value if no students
	}
	// Reuse findTopperPerUniversity to get the overall topper
	toppers := findTopperPerUniversity(students)
	// Assuming we want the overall topper from the map
	var topperO studentStat
	for _, t := range toppers {
		if topperO.finalScore < t.finalScore {
			topperO = t
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
	// Use the Stringer implementation
	fmt.Println("Overall Topper: ", overallTopper)
	fmt.Println("Toppers per University: ")
	for _, topper := range toppersPerUniversity {
		fmt.Println(topper) // Implicitly calls the String() method
	}
}
