package domain

// Class type for defining class domain
type Class struct {
	ID       string
	Students []Student
	Year     int
}

func (c Class) countStudents() int {
	return len(c.Students)
}
