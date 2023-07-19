package main

import "fmt"

func TowerOfHonoi(n int, source, destination, helper string) {
	if n == 1 {
		fmt.Println("move 1th plate to", destination, " from ", source)
		return
	}
	TowerOfHonoi(n-1, source, helper, destination)
	fmt.Println("move ", n, "th plate to ", destination, "from", source)
	TowerOfHonoi(n-1, helper, destination, source)
}

func main() {

	TowerOfHonoi(3, "a", "c", "b")
}
