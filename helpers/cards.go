package helpers

import (
	"strconv"
	"strings"

	"github.com/randy1burrell/toggle-game/data/models"
)

func cardTypes() []string {
	return []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}
}

func cardCodeMap() map[int]string {
	code := make(map[int]string)
	it := 1
	for it <= 13 {
		switch {
		case it == 1:
			code[it] = "ACE"
		case it == 11:
			code[it] = "JACK"
		case it == 12:
			code[it] = "QUEEN"
		case it == 13:
			code[it] = "KING"
		default:
			code[it] = strconv.Itoa(it)
		}
		it = it + 1
	}
	return code
}

func cardCodeArray() []string {
	codeArr := make([]string, 13)

	for i := range codeArr {
		switch {
		case i == 0:
			codeArr[i] = "ACE"
		case i == 10:
			codeArr[i] = "JACK"
		case i == 11:
			codeArr[i] = "QUEEN"
		case i == 12:
			codeArr[i] = "KING"
		default:
			codeArr[i] = strconv.Itoa(i + 1)
		}
	}
	return codeArr
}

func filterCards(cards []models.Card, wanted []string) []models.Card {
	var filtered []models.Card
	for _, val := range cards {
		for _, v := range wanted {
			if strings.ToUpper(v) == val.Code {
				filtered = append(filtered, val)
			}
		}
	}
	return filtered
}

func shuffleCards(wanted []string, word int) []models.Card {
	suit := cardTypes()
	var cards []models.Card
	code := cardCodeMap()
	for _, val := range suit {
		for _, v := range code {
			cd := string([]rune(v)[0]) + string([]rune(val)[0])
			cards = append(cards, models.Card{
				Value: v,
				Suite: val,
				Code:  cd,
			})
		}
	}
	if word > 0 {
		cards = filterCards(cards, wanted)
	}
	return cards
}

func noShuffleCards(wanted []string, word int) []models.Card {
	suit := cardTypes()
	code := cardCodeArray()
	var cards []models.Card
	for _, val := range suit {
		for _, v := range code {
			cd := string([]rune(v)[0]) + string([]rune(val)[0])
			cards = append(cards, models.Card{
				Value: v,
				Suite: val,
				Code:  cd,
			})
		}
	}
	if word > 0 {
		cards = filterCards(cards, wanted)
	}
	return cards
}

// GenerateCards generates a set of cards
func GenerateCards(shuffle bool, wantedStr string) []models.Card {
	var cards []models.Card

	wanted := strings.Split(wantedStr, ",")

	if shuffle {
		cards = shuffleCards(wanted, len(wantedStr))
	} else {
		cards = noShuffleCards(wanted, len(wantedStr))
	}
	return cards
}
