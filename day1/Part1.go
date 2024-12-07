package main

import (
	"advent/utils"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"time"
)

// code qui vient de "https://github.com/patali/advent-of-code-2024" ce qui m'a permis de m'appuyer pour ce début sur
// du Go que je n'ai jamais pu touché avant...
func Day1Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 1 Part1 puzzle: Running")
	sum := 0

	// load file
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day1 input")
	}
	defer file.Close()

	// regex to match numbers
	r, _ := regexp.Compile(`-?\d+`)

	leftList := []int{}
	rightList := []int{}

	// process one line at a time
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := r.FindAllString(scanner.Text(), -1)
		if len(data) == 2 {
			leftNum := utils.StrToInt(data[0])
			rightNum := utils.StrToInt(data[1])
			leftList = append(leftList, leftNum)
			rightList = append(rightList, rightNum)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// sort lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	// calculate distance sum
	for i := 0; i < len(leftList); i++ {
		sum += int(math.Abs(float64(leftList[i]) - float64(rightList[i])))
	}

	fmt.Println("Day 1 Part1 puzzle: Result = ", sum)
}