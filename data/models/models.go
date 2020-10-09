package models

// Deck is a structure to store each Deck
type Deck struct {
	Shuffled  bool   `bson:"shuffled" json:"shuffled"`
	Remaining int    `bson:"remaining" json:"remaining"`
	Cards     []Card `bson:"cards" json:"cards"`
}

type DeckId struct {
	Id        string `bson:"_id" json:"deck_id"`
	Shuffled  bool   `bson:"shuffled" json:"shuffled"`
	Remaining int    `bson:"remaining" json:"remaining"`
	Cards     []Card `bson:"cards" json:"cards"`
}

// Card is a structure for cards
type Card struct {
	Value string `bson:"value" json:"value"`
	Suite string `bson:"suite" json:"suite"`
	Code  string `bson:"code" json:"code"`
}

// DeckMo is a struct for storing deck data
type DeckMo struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

// OpenDeck is a structure to show a Deck and all the cards within
type OpenDeck struct {
	Deck  Deck
	Cards []Card
}

// Testing purposes
type ApiExistsTest struct {
	Endpoint  string
	Method    string
	ResType   int
}

type CreateDeckTest struct {
	Shuffle   string
	Cards     string
	Remaining int
}
