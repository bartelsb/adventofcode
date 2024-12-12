package day3

import (
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Solve(inputFile string) int {
	content, err := os.ReadFile(inputFile)
	check(err)

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(string(content), -1)

	result := 0
	for _, match := range matches {
		if len(match) == 3 {
			x, err := strconv.Atoi(match[1])
			check(err)
			y, err := strconv.Atoi(match[2])
			check(err)
			result = result + (x * y)
		}
	}
	return result
}
