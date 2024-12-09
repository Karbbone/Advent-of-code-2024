package main

import (
	"bufio"
	"day2/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func Part1() {
	fmt.Println("Day 2 Part1 puzzle: Running")
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day1 input")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var intLines [][]int
	for scanner.Scan() {
		intNumbers := utils.TextIntoIntArray(strings.Split(scanner.Text(), " "))
		intLines = append(intLines, intNumbers)
	}
	fmt.Println(intLines)
}
