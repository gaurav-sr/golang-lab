package domain

//Student domain object
type Student struct {
	studentID string
	name      string
	//dateOfBirth time.Time
}

func (s Student) getStudentID() string {
	return s.studentID
}

func (s Student) getStudentName() string {
	return s.name
}

//func (s Student) getDateOfBirth() time.Time {
//	return s.dateOfBirth
//}
