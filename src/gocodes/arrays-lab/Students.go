package main

import "fmt"

// StudentClass structure
type StudentClass struct {
	className string
	names     []string
	count     int
}

func (c *StudentClass) getStudentCount() int {
	return c.count
}

func (c *StudentClass) addStudent(name string) int {
	fmt.Println(name)
	c.names = append(c.names, name)
	c.count++
	return c.count
}
