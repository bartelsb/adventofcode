package day4

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Card struct {
	Number  int
	Winners []int
	Players []int
}

func Solve(inputFile string) int {
	content, err := os.ReadFile(inputFile)
	check(err)

	cards := parseCards(string(content))

	sum := 0
	for _, card := range cards {
		sum += getCardValue(card)
	}
	return sum
}

func parseCards(input string) []Card {
	cards := make([]Card, 0)
	rawCardsList := strings.Split(input, "\n")
	for _, rawCard := range rawCardsList {
		card := new(Card)
		re := regexp.MustCompile(`^Card\s+(\d+):\s+(.*)`)
		matches := re.FindStringSubmatch(rawCard)
		if len(matches) >= 3 {
			number, _ := strconv.Atoi(matches[1])
			card.Number = number
		} else {
			continue
		}
		cardData := strings.Split(matches[2], "|")
		winnerValues := strings.Fields(cardData[0])
		for _, number := range winnerValues {
			winner, _ := strconv.Atoi(number)
			card.Winners = append(card.Winners, winner)
		}
		playerValues := strings.Fields(cardData[1])
		for _, number := range playerValues {
			player, _ := strconv.Atoi(number)
			card.Players = append(card.Players, player)
		}
		cards = append(cards, *card)
	}
	return cards
}

func getCardValue(card Card) int {
	value := 0
	for _, number := range card.Players {
		for _, winner := range card.Winners {
			fmt.Printf("Checking %v and %v\n", number, winner)
			if number == winner {
				if value == 0 {
					value = 1
				} else {
					value = value * 2
				}
			}
		}
	}
	return value
}
