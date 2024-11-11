package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParseStudent(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		want    Student
		wantErr bool
	}{
		{
			name: "valid student",
			line: "John Doe,University A,90,85,88",
			want: Student{
				name:       "John Doe",
				university: "University A",
				grades:     []int{90, 85, 88},
				average:    87.66666666666667,
			},
			wantErr: false,
		},
		{
			name:    "invalid format",
			line:    "John Doe,University A",
			wantErr: true,
		},
		{
			name:    "invalid grade",
			line:    "John Doe,University A,90,invalid,88",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseStudent(tt.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindOverallTopper(t *testing.T) {
	tests := []struct {
		name    string
		students []Student
		want    Student
		wantErr bool
	}{
		{
			name: "normal case",
			students: []Student{
				{name: "John", average: 85},
				{name: "Jane", average: 90},
				{name: "Bob", average: 88},
			},
			want:    Student{name: "Jane", average: 90},
			wantErr: false,
		},
		{
			name:    "empty list",
			students: []Student{},
			want:    Student{},
			wantErr: true,
		},
		{
			name: "single student",
			students: []Student{
				{name: "John", average: 85},
			},
			want:    Student{name: "John", average: 85},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findOverallTopper(tt.students)
			if (err != nil) != tt.wantErr {
				t.Errorf("findOverallTopper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.name != tt.want.name {
				t.Errorf("findOverallTopper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindTopperPerUniversity(t *testing.T) {
	tests := []struct {
		name    string
		students []Student
		want    map[string]Student
		wantErr bool
	}{
		{
			name: "normal case",
			students: []Student{
				{name: "John", university: "Uni A", average: 85},
				{name: "Jane", university: "Uni A", average: 90},
				{name: "Bob", university: "Uni B", average: 88},
			},
			want: map[string]Student{
				"Uni A": {name: "Jane", university: "Uni A", average: 90},
				"Uni B": {name: "Bob", university: "Uni B", average: 88},
			},
			wantErr: false,
		},
		{
			name:    "empty list",
			students: []Student{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "single university",
			students: []Student{
				{name: "John", university: "Uni A", average: 85},
				{name: "Jane", university: "Uni A", average: 90},
			},
			want: map[string]Student{
				"Uni A": {name: "Jane", university: "Uni A", average: 90},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findTopperPerUniversity(tt.students)
			if (err != nil) != tt.wantErr {
				t.Errorf("findTopperPerUniversity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				for uni, student := range tt.want {
					if got[uni].name != student.name || got[uni].average != student.average {
						t.Errorf("findTopperPerUniversity() for %s = %v, want %v", uni, got[uni], student)
					}
				}
			}
		})
	}
} 