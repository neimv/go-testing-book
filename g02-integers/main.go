package main

import "fmt"

// Add takes tow integers and returns the sum of them
func Add(x, y int) int {
	return x + y
}

func main() {
	result := Add(2, 3)
	fmt.Println(result)
}
