package day1

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

var numericStrings = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SumCalibrationValues(inputFile string) int {
	content, err := os.ReadFile(inputFile)
	check(err)

	strContent := string(content)
	arrContent := strings.Split(strContent, "\n")

	sum := 0
	for _, line := range arrContent {
		value := getValueFromLine(line)
		sum = sum + value
	}
	return sum
}

func getValueFromLine(line string) int {
	chars := []rune(line)
	reverseChars := []rune(line)
	slices.Reverse(reverseChars)

	first := getFirstDigit(chars, true)
	last := getFirstDigit(reverseChars, false)

	if first != "" && last != "" {
		value, err := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
		check(err)
		return value
	}
	return 0
}

func getFirstDigit(chars []rune, forwards bool) string {
	for index, character := range chars {
		if unicode.IsDigit(character) {
			return string(character)
		}
		// check if string matches
		for key := range numericStrings {
			if len(chars)-index >= len(key) {
				lineWord := string(chars[index : index+len(key)])
				correctMapWord := ""
				if !forwards {
					mapWord := []rune(key)
					slices.Reverse(mapWord)
					correctMapWord = string(mapWord)
				} else {
					correctMapWord = key
				}
				if lineWord == correctMapWord {
					return numericStrings[key]
				}
			}
		}
	}
	return ""
}
