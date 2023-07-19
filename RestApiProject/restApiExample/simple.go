package main

import "fmt"

func main() {
	var Id int

	fmt.Printf("Enter Student ID: ")
	fmt.Scanln(&Id)
	if Id == 9889 {
		name := "Manju"
		dept := "BCA"
		fee := 15000
		deu := 700
		fmt.Printf("Student found\n+"+
			"Name:", name,
			"\nDepartment: ", dept+
				"\nTotal Fee: "+
				"Deu \n"+
				"Roll no: ", name, dept, fee, deu)
	}
}
