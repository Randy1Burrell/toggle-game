package v1

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gorilla/mux"

	"github.com/randy1burrell/toggle-game/pkg/logger"
)

func TestApiUtils(t *testing.T) {
	// Define router
	r := mux.NewRouter()

	r.NotFoundHandler = &NotFound{}
	r.MethodNotAllowedHandler = &NotAllowed{}
	r.HandleFunc("/deck", CreateDeck).Methods(http.MethodPost)

	// Start the server and listen on the defined port
	go func() {
		logger.Error.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "807"), r))
	}()
	t.Run("Test not found handler", routeNotFound)
	t.Run("Test method not allowed handler", methodNotAllowed)


}

func routeNotFound(t *testing.T) {
	res, err := http.Get("http://test_server:807/d")

	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Error("URL should not be found")
	}
}

func methodNotAllowed(t *testing.T) {
	res, err := http.Get("http://test_server:807/deck")

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Error("Method should be  invalid on URL")
	}
}
