package main

type Student struct {
	id string
	name string
	address string
}

func (s *Student) initialize() []Student{
	s1 := Student{"R101","Gaurav", "Dublin"}
	s2 := Student{"R102","Agni", "Pune"}
	s3 := Student{"R103","Arihant", "Mumbai"}
	s4 := Student{"R104","Tejas", "Nashik"}
	s5 := Student{"R105","Arjun", "Jaipur"}

	students := make([]Student, 0, 5)
	students = append(students, s1, s2, s3, s4, s5)
	return students
}

