package somestructs

import "fmt"

//StudentClass structure
type StudentClass struct {
	// declared as slice
	names []string
}

func (s *StudentClass) addStudent(name string) {
	// use of append function
	s.names = append(s.names, name)
	fmt.Printf("\nAdded %s\n", s.names[len(s.names)-1])
}

func (s *StudentClass) getStudentCount() int {
	return len(s.names)
}
