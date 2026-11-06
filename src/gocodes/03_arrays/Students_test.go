package main

import "testing"

func TestCountStudents(t *testing.T) {
	s := StudentClass{}
	count := s.getStudentCount()
	if count != 0 {
		t.Errorf("Expected %d but found %d", 0, count)
	}
}

func TestAddStudent(t *testing.T) {
	s := StudentClass{"Physics-Class10", nil, 0}
	count := s.count
	position := s.addStudent("GS")
	if position != count+1 {
		t.Errorf("Expected position %d but found %d", count+1, position)
	}
}
