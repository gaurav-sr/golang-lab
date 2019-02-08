package domain

type Class struct {
	Id string
	Students []Student
	Year int
}

func (c Class) countStudents() int {
	return len(c.Students)
}
