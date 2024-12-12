package main

import (
	"fmt"
	"strconv"

	"github.com/bartelsb/adventofcode/2024/day3"
)

func main() {
	result := day3.Solve("./day3/input.txt")
	fmt.Println("Result is " + strconv.Itoa(result))
}
