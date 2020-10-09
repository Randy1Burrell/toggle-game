package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	cards "github.com/randy1burrell/toggle-game/controllers/api/v1"
)

func StartApi(r *mux.Router) *mux.Router {
	// Version my api in case I would like to upgrade in the future
	api := r.PathPrefix("/api/v1").Subrouter()
	// I can always pass api router to as many functions as I like
	api = cardRoute(api)
	// Return the api router to the caller so that is can be used
	return api
}

func cardRoute(r *mux.Router) *mux.Router {
	// Immediately define no found handle
	r.NotFoundHandler = &cards.NotFound{}
	r.MethodNotAllowedHandler = &cards.NotAllowed{}
	// When no query is passed
	r.HandleFunc("/deck", cards.CreateDeck).Methods(http.MethodPost)

	// Get card with uuid
	r.HandleFunc("/deck/{uuid}", cards.OpenDeck).Methods(http.MethodGet)

	// Draw a card or cards from the deck
	r.HandleFunc("/deck/{uuid}/draw", cards.DrawCard).Methods(http.MethodGet).Queries("count", "{[0-9]+}")
	return r
}
