package main

//StudentClass structure
type StudentClass struct {
	names []string
	count int
}

func (c *StudentClass) getStudentCount() int {
	return c.count
}

func (c *StudentClass) addStudent(name string) {

}
