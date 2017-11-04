package main

import "testing"

func TestCountStudents(t *testing.T) {
	s := StudentClass{}
	count := s.getStudentCount()
	if count != -1 {
		t.Errorf("Expected %d but found %d", -1, count)
	}
}

func TestAddStudent(t *testing.T) {
	s := StudentClass{nil, -1}
	count := s.count
	position := s.addStudent("GS")
	if position != count+1 {

	}
}
