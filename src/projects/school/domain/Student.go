package domain

import "time"

//Student domain object
type Student struct {
	StudentID string
	Name      string
	DateOfBirth time.Time
}

func (s *Student) GetStudentID() string {
	return s.StudentID
}

func (s Student) GetStudentName() string {
	return s.Name
}

func (s Student) getDateOfBirth() time.Time {
	return s.DateOfBirth
}

func (s Student) GetAge() int {
	since := time.Since(s.DateOfBirth)
	yr := since/time.Hour/24/365
	return int(yr)
}
