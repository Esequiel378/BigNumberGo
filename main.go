package main

import (
	"fmt"
	"log"
)

func AddNumbers(lhs, rhs string) (string, error) {
	return "", fmt.Errorf("AddNumbers not implemented")
}

func main() {
	result, err := AddNumbers("123 456 789", "11 22 33")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result:", result)
}
