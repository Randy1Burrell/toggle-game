package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/randy1burrell/toggle-game/pkg/application"
	"github.com/randy1burrell/toggle-game/pkg/exithandler"
	"github.com/randy1burrell/toggle-game/pkg/logger"
	"github.com/randy1burrell/toggle-game/routers"
)

func main() {
	// Define router
	r := mux.NewRouter()
	// Send router to function to determine what routes it should define
	app := routers.StartApi(r)

	// Start the server and listen on the defined port
	conf := application.Get()
	cfg := conf.Cfg
	go func() {
		logger.Info.Println("Serving on port", cfg.Port)
		logger.Error.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), app))
	}()

	// Handle Error
	exithandler.Init(func() {
		logger.Info.Println("Server Stopped")
	})
}
