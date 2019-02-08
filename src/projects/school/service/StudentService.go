package main

import (
	"projects/school/domain"
	"time"
)

func main() {
	s := domain.Student{StudentID: "SD101", Name: "Gaurav",
		DateOfBirth: time.Date(1980, 07, 26, 11,11,11, 0 ,time.UTC)}
	println(s.GetStudentID())
	println(s.GetStudentName())
	println(s.GetAge())
}
