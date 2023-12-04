package day2

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CubeCount struct {
	Cubes [3]int // holds 3 values like [red, green blue]
}

type CubeCountBuilder struct {
	red   int
	green int
	blue  int
}

func New() *CubeCountBuilder {
	return &CubeCountBuilder{
		red:   0,
		green: 0,
		blue:  0,
	}
}

func (b *CubeCountBuilder) Build() *CubeCount {
	return &CubeCount{
		Cubes: [3]int{b.red, b.green, b.blue},
	}
}

func (b *CubeCountBuilder) Red(red int) *CubeCountBuilder {
	b.red = red
	return b
}

func (b *CubeCountBuilder) Green(green int) *CubeCountBuilder {
	b.green = green
	return b
}

func (b *CubeCountBuilder) Blue(blue int) *CubeCountBuilder {
	b.blue = blue
	return b
}

type GameResult struct {
	ID     int
	Result []CubeCount
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Solve(inputFile string, redCubes, greenCubes, blueCubes int) int {
	actualCubes := New().Red(redCubes).Green(greenCubes).Blue(blueCubes).Build()

	content, err := os.ReadFile(inputFile)
	check(err)

	strContent := string(content)
	games := strings.Split(strContent, "\n")
	sum := 0
	for _, game := range games {
		gameResult := parseGameData(game)
		sum += isGameValid(actualCubes, gameResult)
	}
	return sum
}

func parseGameData(gameData string) *GameResult {
	gameDataArray := strings.Split(gameData, ": ")
	gameID, err := strconv.Atoi(strings.Split(gameDataArray[0], " ")[1])
	check(err)

	gameResultsArray := strings.Split(gameDataArray[1], ";")
	re := regexp.MustCompile(`\s*(\d+)\s*(blue|red|green)`)
	cubeCounts := []CubeCount{}

	for _, gameResult := range gameResultsArray {
		matches := re.FindAllStringSubmatch(gameResult, -1)
		cubeCount := New()

		for _, match := range matches {
			count, err := strconv.Atoi(match[1])
			check(err)

			color := match[2]
			switch color {
			case "red":
				cubeCount.Red(count)
			case "green":
				cubeCount.Green(count)
			case "blue":
				cubeCount.Blue(count)
			}
		}
		cubeCounts = append(cubeCounts, *cubeCount.Build())
	}
	gameResult := GameResult{gameID, cubeCounts}

	return &gameResult
}

func isGameValid(actual *CubeCount, gameResult *GameResult) int {
	for _, draw := range gameResult.Result {
		fmt.Printf("\nDraw red: %d, draw green: %d, draw blue: %d\n", draw.Cubes[0], draw.Cubes[1], draw.Cubes[2])
		if !(actual.Cubes[0] >= draw.Cubes[0]) || !(actual.Cubes[1] >= draw.Cubes[1]) || !(actual.Cubes[2] >= draw.Cubes[2]) {
			fmt.Printf("Game %d was NOT valid\n", gameResult.ID)
			return 0
		}
	}
	fmt.Printf("Game %d was valid\n", gameResult.ID)
	return gameResult.ID
}
