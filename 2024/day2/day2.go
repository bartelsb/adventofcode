package day2

import (
	"os"
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

	reports := [][]int{}
	for _, report := range strings.Split(string(content), "\n") {
		reportArr := []int{}
		for _, level := range strings.Split(report, " ") {
			numLevel, err := strconv.Atoi(level)
			check(err)
			reportArr = append(reportArr, numLevel)
		}
		reports = append(reports, reportArr)
	}

	safeReports := 0
	for _, report := range reports {
		safe := true
		direction := true
		issues := 0
		if report[0] != report[1] {
			direction = checkDirection(report[0], report[1])
		}
		for i := 0; i < len(report)-1; i++ {
			safe = checkLevelDiff(report[i], report[i+1], direction)
			if !safe {
				issues++
				if issues > 1 {
					break
				}
				safe = true
			}
		}
		if safe {
			safeReports++
		}
	}

	return safeReports
}

func checkLevelDiff(level1, level2 int, direction bool) bool {
	diff := level1 - level2
	if direction {
		return (diff >= -3) && (diff <= -1)
	} else {
		return (diff >= 1) && (diff <= 3)
	}
}

func checkDirection(level1, level2 int) bool {
	return level1 < level2
}
