package day1

import (
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Solve(inputFile string) int {
	content, err := os.ReadFile(inputFile)
	check(err)

	leftArr, rightArr := parseInput(content)
	sort.Ints(leftArr)
	sort.Ints(rightArr)

	totalDistance := 0
	for i := 0; i < len(leftArr); i++ {
		distance := math.Abs((float64)(leftArr[i] - rightArr[i]))
		totalDistance += int(distance)
	}
	return totalDistance
}

func parseInput(content []byte) ([]int, []int) {
	var leftArr []int
	var rightArr []int
	re := regexp.MustCompile(`\s*(\d+)\s+(\d+)\s*`)

	for _, line := range strings.Split(string(content), "\n") {
		if len(line) == 0 {
			continue
		}
		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			continue
		}
		leftnum, err := strconv.Atoi(matches[1])
		check(err)
		rightnum, err := strconv.Atoi(matches[2])
		check(err)
		leftArr = append(leftArr, leftnum)
		rightArr = append(rightArr, rightnum)
	}
	return leftArr, rightArr
}
