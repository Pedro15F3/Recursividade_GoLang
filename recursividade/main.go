package main

import (
	"fmt"
	"strconv"
)

func mDecRecursive(s string) int {
	if len(s) == 1 {
		fmt.Printf("Mdec(<%s>) = %s\n", s, s)
		num, _ := strconv.Atoi(s)
		return num
	} else {
		prefix := s[:len(s)-1]
		lastDigit := s[len(s)-1]

		subResult := mDecRecursive(prefix)
		lastDigitValue, _ := strconv.Atoi(string(lastDigit))
		result := 10*subResult + lastDigitValue

		fmt.Printf("Mdec(<%s> '%s') = 10 * Mdec(<%s>) + %s = %d\n", prefix, string(lastDigit), prefix, string(lastDigit), result)
		return result
	}
}

func main() {
	var number string
	fmt.Print("Digite um número: ")
	fmt.Scanln(&number)

	result := mDecRecursive(number)
	fmt.Printf("Final result: %d\n", result)
}
