package structs

import "testing"

func TestCountStudents(t *testing.T) {
	s := StudentClass{}
	count := s.getStudentCount()
	if count != 0 {
		t.Errorf("Expected %d but found %d", -1, count)
	}
}

func TestAddStudent(t *testing.T) {
	s := StudentClass{}
	beforeAddCount := len(s.names) - 1
	s.addStudent("GS")
	afterAddCount := len(s.names) - 1
	if afterAddCount != beforeAddCount+1 {
		t.Errorf("Expected %d, Actual %d", beforeAddCount+1, afterAddCount)
	}
}
