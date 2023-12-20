package main

import (
	"fmt"
	"strconv"

	"github.com/bartelsb/adventofcode/2023/day4"
)

func main() {
	result := day4.Solve("./day4/input.txt")
	fmt.Println("Result is " + strconv.Itoa(result))
}
