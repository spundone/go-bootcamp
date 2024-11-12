# Student Grading System in Go

A Go program that processes student grades from a CSV file and identifies top performers overall and per university.

## Features

- Reads student data from CSV files
- Calculates average grades per student
- Finds the overall topper across all universities
- Identifies toppers for each university
- Implements error handling for file operations and data parsing
- Uses Go's standard library for CSV processing

## Implementation Details

The program uses a `Student` struct to store student information:
```go
type Student struct {
    name       string
    university string
    grades     []int
    average    float64
}
```

### Core Functions

1. `readAndParseCSV`: Reads student data from a CSV file
2. `parseStudent`: Parses individual student records
3. `findOverallTopper`: Identifies the student with highest average
4. `findTopperPerUniversity`: Groups students by university and finds toppers

## Usage

1. Prepare a CSV file (e.g., `students.csv`) with the following format:
```
Name,University,Grade1,Grade2,Grade3,...
John Doe,University A,90,85,88
Jane Smith,University B,92,88,90
```

2. Run the program:
```bash
go run sgp.go
```

3. The program will output:
- Overall topper across all universities
- Toppers for each individual university

## Example Output

```
Overall topper: Jane Smith from University B with average 90.00
Topper for University A: John Doe with average 87.67
Topper for University B: Jane Smith with average 90.00
```

## Code Structure

The main implementation can be found in `sgp.go`:
```go:student-grading-problem/sgp.go
startLine: 11
endLine: 75
```

Tests are implemented in `sgp_test.go`:
```go:student-grading-problem/sgp_test.go
startLine: 9
endLine: 98
```

## Testing

Run the tests using:
```bash
go test
```

The test suite includes:
- Validation of student parsing
- Testing of topper identification
- Error handling verification
- Edge cases for empty input and invalid data

## Prerequisites

- Go 1.21 or higher
- Basic understanding of CSV file format

## Error Handling

The program handles various error cases:
- Invalid CSV format
- Missing or malformed files
- Invalid grade values
- Empty student lists

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License.

## Acknowledgments

- Inspired by [one2nc/student-grading-go](https://github.com/one2nc/student-grading-go)
- Part of the [one2n.io Go Bootcamp](https://one2n.io/go-bootcamp) course
