package main

import (
	"bufio"
	"fmt"
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

	zeroCount := 0
	currentVal := startNum
	for _, input := range inputs {
		currentVal, err = nextVal(currentVal, input)
		if err != nil {
			fmt.Printf("Error getting next val: %v", err)
			os.Exit(1)
		}
		if currentVal == 0 {
			zeroCount++
		}
	}
	fmt.Println("The actual password is:", zeroCount)

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

func nextVal(currentVal int, input string) (int, error) {
	absVal, err := strconv.Atoi(input[1:])
	if err != nil {
		return 0, err
	}
	if strings.HasPrefix(input, "L") {
		return (currentVal - absVal) % 100, nil
	} else {
		return (currentVal + absVal) % 100, nil
	}
}
