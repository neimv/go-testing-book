package main

import "fmt"

const repeatCount = 5

func Repeat(charater string) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated = repeated + charater
	}

	return repeated
}

func main() {
	fmt.Println("vim-go")
}
