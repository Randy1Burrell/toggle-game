package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/gorilla/mux"

	"github.com/randy1burrell/toggle-game/data/models"
	"github.com/randy1burrell/toggle-game/pkg/logger"
)

var host = "http://test_server:808/"

func TestCardDeckApi(t *testing.T) {
	// Define router
	r := mux.NewRouter()

	r.HandleFunc("/deck", CreateDeck).Methods(http.MethodPost)
	r.HandleFunc("/deck/{uuid}", OpenDeck).Methods(http.MethodGet)
	r.HandleFunc("/deck/{uuid}/draw", DrawCard).Methods(http.MethodGet).Queries("count", "{[0-9]+}")

	// Start the server and listen on the defined port
	go func() {
		logger.Error.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "808"), r))
	}()

	t.Run("Test create deck success", createDeckSuccess)
	t.Run("Test create deck bad request", createDeckBadRequest)
	t.Run("Test create deck ignore bad forms", createDeckIgnoreUnrecFormData)
	t.Run("Test open valid deck", openValidDeck)
	t.Run("Test open invalid deck not found", openDeckNotFound)
	t.Run("Test draw card with invalid count", drawCardInvalidCount)
	t.Run("Test draw card when count is too large", drawCardCountTooLarge)
	t.Run("Test draw card when count is too small", drawCardCountTooSmall)
	t.Run("Test draw card when count is valid", drawCardValidCount)
	t.Run("Test draw card when uuid is invalid", drawCardInvalidUuid)
	t.Run("Test draw card when deck is empty", drawCardFromEmptyDeck)
}

func createDeckSuccess(t *testing.T) {
	testData := []models.CreateDeckTest{
		{"", "", 52},
		{"true", "", 52},
		{"false", "", 52},
		{"true", "kdfdfk", 0},
		{"false", "kdfdfk", 0},
		{"", "hd,df", 0},
		{"true", "hd,df", 0},
		{"false", "hd,df", 0},
		{"", "AS,KD,AC,2C,KH", 5},
		{"true", "AS,KD,AC,2C,KH", 5},
		{"false", "AS,KD,AC,2C,KH", 5},
	}


	for _, val := range testData {
		res, err := http.PostForm(host + "deck", url.Values{"shuffle": {val.Shuffle}, "cards": {val.Cards}})
		if err != nil {
			t.Error(err)
		}
		body, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		deck := &models.DeckMo{}
		json.Unmarshal(body, &deck)
		if len(deck.DeckID) <= 0 {
			t.Error("deck_id is wrong")
		}
		if val.Remaining != deck.Remaining {
			t.Error("Remaining is wrong")
		}
	}
}

func createDeckBadRequest(t *testing.T) {
	testData := []models.CreateDeckTest{
		{"n", "sdds", 0},
		{"tru", "", 0},
		{"fals", "", 0},
		{"tre", "kdfdfk", 0},
		{"flse", "kdfdfk", 0},
		{"s", "hd,df", 0},
		{"tr", "hd,df", 0},
		{"fls", "AS,KD,AC,2C,KH", 0},
	}

	for _, val := range testData {
		res, err := http.PostForm(host + "deck", url.Values{"shuffle": {val.Shuffle}, "cards": {val.Cards}})
		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != http.StatusBadRequest {
			t.Error("Server should return a bad request message")
		}
	}
}

func createDeckIgnoreUnrecFormData(t *testing.T) {
	testData := []models.CreateDeckTest{
		{"", "", 52},
		{"true", "", 52},
		{"false", "", 52},
		{"true", "kdfdfk", 0},
		{"false", "kdfdfk", 0},
		{"", "hd,df", 0},
		{"true", "hd,df", 0},
		{"false", "hd,df", 0},
		{"", "AS,KD,AC,2C,KH", 5},
		{"true", "AS,KD,AC,2C,KH", 5},
		{"false", "AS,KD,AC,2C,KH", 5},
	}


	for _, val := range testData {
		res, err := http.PostForm(host + "deck", url.Values{"shuffl": {val.Shuffle}, "cards": {val.Cards}})
		if err != nil {
			t.Error(err)
		}
		body, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		deck := &models.DeckMo{}
		json.Unmarshal(body, &deck)
		if len(deck.DeckID) <= 0 {
			t.Error("deck_id is wrong")
		}
		if val.Remaining != deck.Remaining {
			t.Error("Remaining is wrong")
		}
	}
}

func openValidDeck(t *testing.T) {
	testData := []models.CreateDeckTest{
		{"", "", 52},
		{"true", "", 52},
		{"false", "", 52},
		{"true", "kdfdfk", 0},
		{"false", "kdfdfk", 0},
		{"", "hd,df", 0},
		{"true", "hd,df", 0},
		{"false", "hd,df", 0},
		{"", "AS,KD,AC,2C,KH", 5},
		{"true", "AS,KD,AC,2C,KH", 5},
		{"false", "AS,KD,AC,2C,KH", 5},
	}


	for _, val := range testData {
		res, err := http.PostForm(host + "deck", url.Values{"shuffl": {val.Shuffle}, "cards": {val.Cards}})
		if err != nil {
			t.Error(err)
		}
		body, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		deck := &models.DeckMo{}
		json.Unmarshal(body, &deck)
		if len(deck.DeckID) <= 0 {
			t.Error("deck_id is wrong")
		}
		if val.Remaining != deck.Remaining {
			t.Error("Remaining is wrong")
		}

		res, err = http.Get(host + "deck/" + deck.DeckID)

		if err != nil {
			t.Error(err)
		}

		entireDeck := &models.DeckId{}
		body, err = ioutil.ReadAll(res.Body)
		res.Body.Close()
		json.Unmarshal(body, &entireDeck)

		if entireDeck.Id != deck.DeckID {
			t.Error("Wrong deck returned")
		}

		if entireDeck.Shuffled != deck.Shuffled {
			t.Error("Shuffle is wrong")
		}

		remaining := len(entireDeck.Cards)
		if remaining != deck.Remaining || remaining != entireDeck.Remaining {
			t.Error("Cards remaining in deck showing wrong value")
		}
	}
}

func openDeckNotFound(t *testing.T) {
	testData := []models.CreateDeckTest{
		{"", "", 52},
		{"true", "", 52},
		{"false", "", 52},
		{"true", "kdfdfk", 0},
		{"false", "kdfdfk", 0},
		{"", "hd,df", 0},
		{"true", "hd,df", 0},
		{"false", "hd,df", 0},
		{"", "AS,KD,AC,2C,KH", 5},
		{"true", "AS,KD,AC,2C,KH", 5},
		{"false", "AS,KD,AC,2C,KH", 5},
	}


	for _, val := range testData {
		res, err := http.PostForm(host + "deck", url.Values{"shuffl": {val.Shuffle}, "cards": {val.Cards}})
		if err != nil {
			t.Error(err)
		}
		body, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		deck := &models.DeckMo{}
		json.Unmarshal(body, &deck)
		if len(deck.DeckID) <= 0 {
			t.Error("deck_id is wrong")
		}
		if val.Remaining != deck.Remaining {
			t.Error("Remaining is wrong")
		}

		res, err = http.Get(host + "deck/" + deck.DeckID + "kjsh")

		if err != nil {
			t.Error(err)
		}


		if res.StatusCode != http.StatusNotFound {
			t.Error("Wrong data found")
		}
	}
}

func drawCardInvalidCount(t *testing.T) {
	res, err := http.PostForm(host + "deck", url.Values{"shuffl": {"true"}, "cards": {""}})
	if err != nil {
		t.Error(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	deck := &models.DeckMo{}
	json.Unmarshal(body, &deck)
	if len(deck.DeckID) <= 0 {
		t.Error("deck_id is wrong")
	}
	res, err = http.Get(host + "deck/" + deck.DeckID + "/draw?count=" + deck.DeckID)
	if res.StatusCode != http.StatusBadRequest {
		t.Error("Request should be a bad request")
	}
}

func drawCardCountTooLarge(t *testing.T) {
	res, err := http.PostForm(host + "deck", url.Values{"shuffl": {"true"}, "cards": {""}})
	if err != nil {
		t.Error(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	deck := &models.DeckMo{}
	json.Unmarshal(body, &deck)
	if len(deck.DeckID) <= 0 {
		t.Error("deck_id is wrong")
	}
	res, err = http.Get(host + "deck/" + deck.DeckID + "/draw?count=53")
	if res.StatusCode != http.StatusBadRequest {
		t.Error("Request should be a bad request")
	}
}

func drawCardCountTooSmall(t *testing.T) {
	res, err := http.PostForm(host + "deck", url.Values{"shuffl": {"true"}, "cards": {""}})
	if err != nil {
		t.Error(err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	deck := &models.DeckMo{}
	json.Unmarshal(body, &deck)
	if len(deck.DeckID) <= 0 {
		t.Error("deck_id is wrong")
	}
	res, err = http.Get(host + "deck/" + deck.DeckID + "/draw?count=0")
	if res.StatusCode != http.StatusBadRequest {
		t.Error("Request should be a bad request")
	}
}

func drawCardValidCount(t *testing.T) {
	for i := 1; i < 10; i++ {
		res, err := http.PostForm(host + "deck", url.Values{})
		if err != nil {
			t.Error(err)
		}
		body, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		deck := &models.DeckMo{}
		json.Unmarshal(body, &deck)
		if len(deck.DeckID) <= 0 {
			t.Error("deck_id is wrong")
		}
		count := strconv.Itoa(i)
		res, err = http.Get(host + "deck/" + deck.DeckID + "/draw?count=" + count)
		if err != nil {
			t.Error(err)
		}

		var cards = []models.Card{}
		body, err = ioutil.ReadAll(res.Body)
		res.Body.Close()
		json.Unmarshal(body, &cards)

		if len(cards) != i {
			t.Error("Not getting correct amount of cards from api")
		}
	}
}

func drawCardInvalidUuid(t *testing.T) {
	for i := 1; i < 10; i++ {
		res, err := http.PostForm(host + "deck", url.Values{})
		if err != nil {
			t.Error(err)
		}
		body, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		deck := &models.DeckMo{}
		json.Unmarshal(body, &deck)
		if len(deck.DeckID) <= 0 {
			t.Error("deck_id is wrong")
		}
		count := strconv.Itoa(i)
		res, err = http.Get(host + "deck/" + deck.DeckID + count + "/draw?count=" + count)
		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != http.StatusBadRequest {
			t.Error("Returning cards that should not be returned")
		}
	}
}

func drawCardFromEmptyDeck(t *testing.T) {
	for i := 1; i < 10; i++ {
		res, err := http.PostForm(host + "deck", url.Values{})
		if err != nil {
			t.Error(err)
		}
		body, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		deck := &models.DeckMo{}
		json.Unmarshal(body, &deck)
		if len(deck.DeckID) <= 0 {
			t.Error("deck_id is wrong")
		}
		count := strconv.Itoa(deck.Remaining)
		res, err = http.Get(host + "deck/" + deck.DeckID + "/draw?count=" + count)
		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != http.StatusOK {
			t.Error("Should return correct amount of cards")
		}

		res, err = http.Get(host + "deck/" + deck.DeckID + "/draw?count=" + count)
		if err != nil {
			t.Error(err)
		}

		body, _ = ioutil.ReadAll(res.Body)
		res.Body.Close()
		cards := &models.Card{}
		json.Unmarshal(body, &cards)

		if res.StatusCode != http.StatusNoContent {
			logger.Info.Println(cards)
			t.Error("Should return correct amount of cards")
		}
	}
}
