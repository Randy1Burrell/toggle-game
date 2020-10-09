package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/randy1burrell/toggle-game/data/models"
)

var host = "http://test_server:8080/"

func TestCreateDeckSuccess(t *testing.T) {
	go func() {
		main()
	}()

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

	t.Run("Test create deck", func(t *testing.T) {
		for _, val := range testData {
			res, err := http.PostForm(host+`api/v1/deck`, url.Values{"shuffle": {val.Shuffle}, "cards": {val.Cards}})
			if err != nil {
				t.Error(err)
			}
			body, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			deck := &models.DeckMo{}
			json.Unmarshal(body, &deck)
			if len(deck.DeckID) <= 0 {
				t.Error("UUID NOT VALID")
			}
			if val.Remaining != deck.Remaining {
				t.Error("Remaining is wrong")
			}
		}
	})
}
