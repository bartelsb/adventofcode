package main

import (
	"fmt"

	"github.com/bartelsb/adventofcode/2023/day1"
)

func main() {
	sum := day1.SumCalibrationValues("./day1/input.txt")
	result := fmt.Sprintf("Result is %d", sum)
	fmt.Println(result)
}
