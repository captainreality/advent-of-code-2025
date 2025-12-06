package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	startNum = 50
)

func main() {
	fmt.Println("Welcome to the secret entrance!")
	inputs, err := readInput("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v", err)
		os.Exit(1)
	}

	currentVal := startNum
	totalCrossings := 0
	for _, input := range inputs {
		crossings := 0
		currentVal, crossings, err = nextVal(currentVal, input)
		if err != nil {
			fmt.Printf("Error getting next val: %v", err)
			os.Exit(1)
		}
		totalCrossings += crossings
	}
	fmt.Println("The actual password is:", totalCrossings)
}

func readInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func nextVal(currentVal int, input string) (nextVal int, crossings int, err error) {
	inVal, err := strconv.Atoi(input[1:])
	if err != nil {
		return 0, 0, err
	}
	inter := 0
	if strings.HasPrefix(input, "L") {
		inter = currentVal - inVal
	} else {
		inter = currentVal + inVal
	}

	crossings = int(math.Abs(float64(inter / 100)))
	if currentVal != 0 && inter <= 0 {
		crossings++
	}

	nextVal = inter % 100
	// At this point, nextVal could be negative
	if nextVal < 0 {
		nextVal = 100 + nextVal
	}

	return nextVal, crossings, nil
}
