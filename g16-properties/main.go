package main

import (
	"fmt"
	"strings"
)

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for i := 0; i < arabic; i++ {
		if i == 4 {
			result.WriteString("IV")
			break
		}

		result.WriteString("I")
	}

	return result.String()
}

func main() {
	fmt.Println("vim-go")
}
