package pokdeng

import (
	"math/rand"
	"strconv"
)

var Deck = []int{}

func Pokdeng(p1 int, p2 int) string {

	Deck = CreateCard()
	// p1 = calScore(RandomCard())
	// p2 = calScore(RandomCard())

	if p1 > p2 {
		return "Player 1 is winner"
	} else if p2 > p1 {
		return "Player 2 is winner"
	} else {
		return "draw"
	}
}

func CalScore(cards []int) int {

	score := 0

	for i := 0; i < len(cards); i++ {
		score += cards[i]
	}

	scoreString := strconv.Itoa(score)
	if len(scoreString) > 1 {
		scoreString = scoreString[1:2]
	}
	intVar, _ := strconv.Atoi(scoreString)

	return intVar
}

func CreateCard() []int {
	deck := []int{}

	for i := 1; i <= 10; i++ {
		for j := 0; j < 4; j++ {
			deck = append(deck, i)
		}
	}
	return deck
}

func RandomCard() []int {

	holdCard := []int{}

	for i := 0; i < 2; i++ {
		randomIndex := rand.Intn(len(Deck))
		randomCard := Deck[randomIndex]
		Deck = RemoveIndex(Deck, randomIndex)
		holdCard = append(holdCard, randomCard)
	}

	return holdCard
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
