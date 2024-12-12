package main

import (
	"fmt"
	"strconv"

	"github.com/bartelsb/adventofcode/2024/day2"
)

func main() {
	result := day2.Solve("./day2/input.txt")
	fmt.Println("Result is " + strconv.Itoa(result))
}
