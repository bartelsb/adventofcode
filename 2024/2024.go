package main

import (
	"fmt"
	"strconv"

	"github.com/bartelsb/adventofcode/2024/day1"
)

func main() {
	result := day1.Solve("./day1/input.txt")
	fmt.Println("Result is " + strconv.Itoa(result))
}
