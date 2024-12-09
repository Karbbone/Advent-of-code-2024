package utils

import (
	"log"
	"strconv"
)

func TextIntoIntArray(array []string) []int {
	var intNumbers []int
	for _, strNumber := range array {
		intNumber, err := strconv.Atoi(strNumber)
		if err != nil {
			log.Fatal("Failed to convert string to int")
		}
		intNumbers = append(intNumbers, intNumber)
	}
	return intNumbers
}
