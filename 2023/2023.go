package main

import (
	"fmt"
	"strconv"

	"github.com/bartelsb/adventofcode/2023/day2"
)

func main() {
	result := day2.Solve("./day2/input.txt", 12, 13, 14)
	fmt.Println("Result is " + strconv.Itoa(result))
}
