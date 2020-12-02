package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	var parsedInput = parseInput(string(content))
	var multiplication = findYearCalculation(parsedInput)

	fmt.Println(multiplication)
}

func parseInput(values string) []int {
	var sliceInt []int
	var splittedString = strings.Split(values, "\n")

	for i := 0; i < len(splittedString); i++ {
		currentValue, _ := strconv.Atoi(splittedString[i])
		sliceInt = append(sliceInt, currentValue)
	}

	return sliceInt
}

func findYearCalculation(values []int) int {
	var multiplication int = 0

	for i := 0; i < (len(values) - 1); i++ {
		for j := 1; j < (len(values) - i); j++ {
			if values[i]+values[j+i] == 2020 {
				multiplication = values[i] * values[j+i]
				break
			}
		}
	}

	return multiplication
}
