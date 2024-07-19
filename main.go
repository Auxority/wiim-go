package main

import (
	"net/http"

	"github.com/Auxority/wiim-go/api"
	"github.com/Auxority/wiim-go/device"
	"github.com/Auxority/wiim-go/wiim"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	address         = ":8080"
	TogglePlayRoute = "/toggle-play"
	StatusRoute     = "/status"
)

func registerRoutes(router *api.Router) {
	http.HandleFunc(TogglePlayRoute, router.TogglePlay)
	http.HandleFunc(StatusRoute, router.GetStatus)
}

func startRouter() {
	log.Info().Str("address", address).Msg("Starting server")

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	ip := device.Find()
	if ip == "" {
		log.Fatal().Msg("No WiiM device found")
	}

	wiimAPI := wiim.New(ip)
	router := api.New(wiimAPI)

	registerRoutes(router)
	startRouter()

	// test, err := wiimAPI.GetStatus()
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("Failed to get status")
	// }

	// // Marshal the struct to JSON
	// jsonData, err := json.MarshalIndent(test, "", "  ")
	// if err != nil {
	// 	fmt.Println("Error marshaling to JSON:", err)
	// 	return
	// }

	// // Print the JSON string
	// fmt.Println(string(jsonData))
}
