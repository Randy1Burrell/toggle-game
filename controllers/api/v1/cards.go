package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/randy1burrell/toggle-game/data/database"
	"github.com/randy1burrell/toggle-game/data/models"
	"github.com/randy1burrell/toggle-game/helpers"
	"github.com/randy1burrell/toggle-game/pkg/logger"
)

var message = `{"message": "%s"}`

// CreateDeck creates a deck
func CreateDeck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// We only care about the shuffle form value
	shuf := r.FormValue("shuffle")
	wantedCards := r.FormValue("cards")

	if len(shuf) == 0 {
		shuf = "false"
	}

	shuffle, err := strconv.ParseBool(shuf)
	if err != nil {
		logger.Info.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf(message, `Shuffle must be a boolean of the form true/false`)
		json.NewEncoder(w).Encode(msg)
		return
	}

	cards := helpers.GenerateCards(shuffle, wantedCards)
	deck := models.Deck{
		Shuffled:  shuffle,
		Remaining: len(cards),
		Cards:     cards,
	}

	deckID, err := database.Insert(&deck)

	if err != nil {
		logger.Info.Println(err.Error())
		msg := fmt.Sprintf(message, `Shuffle must be a boolean of the form true/false`)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusCreated)
	createdDeck := models.DeckMo{
		DeckID:    fmt.Sprintf("%v", deckID),
		Shuffled:  shuffle,
		Remaining: len(cards),
	}

	json.NewEncoder(w).Encode(createdDeck)
}

// OpenDeck opens a deck with the remaining cards
func OpenDeck(w http.ResponseWriter, r *http.Request) {
	// Get parameters from the path
	pathParams := mux.Vars(r)

	// Let's set the some headers because we have to use it regardless of the outcome
	w.Header().Set("Content-Type", "application/json")

	/***
	 * Let's check for the uuid in the path param, we'll need it to
	 * get the card deck from the db, if not found the something is
	 * wrong
	 */
	if uuid, found := pathParams["uuid"]; found {
		if deck, err := database.FindOne(uuid); err == nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(deck)
			return
		}
		// Message for resources not found
		response := fmt.Sprintf(`{"message": "Resource with %s not found"}`, uuid)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	// If no uuid is passed for the path param then this was a bad request
	w.WriteHeader(http.StatusBadRequest)
	response := `{"message": "uuid is a required path parameter"}`
	json.NewEncoder(w).Encode(response)
}

// drawCard retrieves cards from a deck
func DrawCard(w http.ResponseWriter, r *http.Request) {
	// We need count query so if it is not found then lets return an error
	query := r.URL.Query()

	cards := query["count"]
	var countStr string
	if len(cards) > 0 {
		countStr = cards[0]
	}

	// Get parameters from the path
	pathParams := mux.Vars(r)

	// Let's set the some headers because we have to use it regardless of the outcome
	w.Header().Set("Content-Type", "application/json")

	// Count should always be between 1 and 52 inclusive
	if count, err := strconv.Atoi(countStr); err == nil && count <= 52 && count >= 1 {
		if uuid, found := pathParams["uuid"]; found {
			if deck, err := database.DrawCard(uuid, count); err == nil {
				logger.Info.Println(len(deck))
				if len(deck) <= 0 {
					w.WriteHeader(http.StatusNoContent)
				} else {
					w.WriteHeader(http.StatusOK)
				}
				json.NewEncoder(w).Encode(deck)
				return
			}
			response := fmt.Sprintf(`{"message": "Requested resource with %s uuid not found"}`, uuid)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	response := fmt.Sprintf(`{"message": "count query is required to be an integer that is less than 53 and greater than 0"}`)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}
