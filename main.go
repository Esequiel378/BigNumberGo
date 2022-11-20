package main

import (
	"fmt"
	"log"
)

// AddNumbers takse two string params containing M numbers
// separated by spaces and returns sum of the pairs.
//
// + The input parameters should have the same count of numbers.
// + The numbers may include decimal places.
// + The numbers can be arbitrarily long, e.g. 1000+ digits.
//
// Examples:
//  >> AddNumbers("123 456 789", "11 22 33")
//  "134 478 822"
//
//  >> AddNumbers("123456789012345678901 23456789", "12345678 234567890123456789012")
//  "123456789012358024579 234567890123480245801"
//
//  >> AddNumbers("1234567.8901 2.345", "12.34 2345678901.2")
//  "1234580.2301 2345678903.545"
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
