package main

import "fmt"

func fib(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	fmt.Println("n= ", n)
	return fib(n-1) + fib(n-2)
}
func main() {
	fmt.Println(fib(4))
}
