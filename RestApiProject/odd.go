package main

import "fmt"

func main() {

	var a int

	fmt.Println("Enter the number to find the sum and display ")
	fmt.Scanln(&a)
	count := 0
	sum := 0
	for i := 1; i <= 100; i++ {
		if count == a {
			break
		} else if i%2 != 0 {
			count += 1
			sum += i
			fmt.Print(i, " ")
		}
	}
	fmt.Println("Sum of given Odd number is :", sum)
}
