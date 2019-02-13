package main

import "fmt"

func main() {
	fmt.Println("Hello Arrays..")
	format1()
	format2()
	format3()
	format4()
}

func format1() {
	names := [4]string{"Kanpur", "Nasik", "Pune", "Dublin"}
	for idx, name := range names {
		fmt.Printf("%d %s\n", idx, name)
	}
}

func format2() {
	cities := [...]string{"Kanpur", "Pune", "Nasik", "Dublin"}
	for _, city := range cities {
		fmt.Println(city)
	}
}

func format3() {
	var years = [4]int{1999, 2002, 2004, 2014}
	for _, year := range years {
		fmt.Println(year)
	}
}

func format4() {
	doubleDim := [...][3]int{{1, 2, 3},
		{4, 5, 6}}
	for _, row := range doubleDim {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println("")
	}
}
