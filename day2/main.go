package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	wordList, searchedLetters, minAmounts, maxAmounts := parseInput(content)
	var passedPasswordCount = checkAllPasswords(wordList, searchedLetters, minAmounts, maxAmounts)

	fmt.Println("---Amount of passed passwords is:---", passedPasswordCount)
}

func parseInput(input []byte) ([]string, []rune, []int, []int) {
	var stringInput = string(input)
	var slicesOfLines = strings.Split(stringInput, "\n")
	var occurences []string
	var wordList []string
	var searchedLetters []rune
	var occurencesMinMax []string
	var occurenceAmountsMinSanitized []int
	var occurenceAmountsMaxSanitized []int

	for _, row := range slicesOfLines {
		var splitRowsByColon = strings.Split(row, ": ")

		occurences = append(occurences, splitRowsByColon[0])
		wordList = append(wordList, splitRowsByColon[1])
	}

	for _, occurence := range occurences {
		var sliceBySpace = strings.Split(occurence, " ")
		var searchedLetterRune = []rune(sliceBySpace[1])

		searchedLetters = append(searchedLetters, searchedLetterRune[0])
		occurencesMinMax = append(occurencesMinMax, sliceBySpace[0])
	}

	for _, occurenceMinMax := range occurencesMinMax {
		var occurenceMinMaxSlice = strings.Split(occurenceMinMax, "-")
		var minAmountInt, _ = strconv.Atoi(occurenceMinMaxSlice[0])
		var maxAmountInt, _ = strconv.Atoi(occurenceMinMaxSlice[1])

		occurenceAmountsMinSanitized = append(occurenceAmountsMinSanitized, minAmountInt)
		occurenceAmountsMaxSanitized = append(occurenceAmountsMaxSanitized, maxAmountInt)
	}

	return wordList, searchedLetters, occurenceAmountsMinSanitized, occurenceAmountsMaxSanitized
}

func checkOccurences(word string, searchedLetter rune) int {
	var occurenceCount int

	for _, letter := range word {
		if letter == searchedLetter {
			occurenceCount++
		}
	}

	return occurenceCount
}

func checkPassword(password string, letter rune, minAmount int, maxAmount int) bool {
	var amountOfOccurences = checkOccurences(password, letter)

	if amountOfOccurences >= minAmount && amountOfOccurences <= maxAmount {
		return true
	}

	return false
}

func checkAllPasswords(passwordsList []string, letterList []rune, minAmountList []int, maxAmountList []int) int {
	fmt.Println("Counting password in")
	var passedPasswordCount int

	for i, password := range passwordsList {
		if checkPassword(password, letterList[i], minAmountList[i], maxAmountList[i]) {
			passedPasswordCount++
		}
	}

	return passedPasswordCount
}
