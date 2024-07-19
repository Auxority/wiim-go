package api

import (
	"encoding/json"
	"net/http"

	"github.com/Auxority/wiim-go/wiim"
	"github.com/rs/zerolog/log"
)

type Router struct {
	wiimAPI *wiim.API
}

func convertToAPIStatusResponse(wiimStatus *wiim.StatusResponse) *StatusResponse {
	return &StatusResponse{
		Album:             string(wiimStatus.Album),
		Artist:            string(wiimStatus.Artist),
		Status:            wiimStatus.Status.String(),
		Title:             string(wiimStatus.Title),
		Vendor:            wiimStatus.Vendor,
		Channel:           wiimStatus.Channel.String(),
		EqPreset:          int(wiimStatus.EqPreset),
		IsMuted:           bool(wiimStatus.IsMuted),
		RepeatMode:        wiimStatus.LoopMode.String(),
		PlaybackMode:      wiimStatus.PlaybackMode.String(),
		TotalTracks:       int(wiimStatus.TotalTracks),
		TrackDurationMs:   int(wiimStatus.TrackDurationMs),
		CurrentTrackIndex: int(wiimStatus.CurrentTrackIndex),
		TrackProgressMs:   int(wiimStatus.TrackProgressMs),
		Volume:            int(wiimStatus.Volume),
	}
}

func New(wiimAPI *wiim.API) *Router {
	return &Router{
		wiimAPI: wiimAPI,
	}
}

func (a *Router) TogglePlay(w http.ResponseWriter, r *http.Request) {
	err := a.wiimAPI.TogglePlay()
	if err != nil {
		log.Error().Err(err).Msg("Failed to toggle play")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *Router) GetStatus(w http.ResponseWriter, r *http.Request) {
	wiimStatus, err := a.wiimAPI.GetStatus()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get status")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	status := convertToAPIStatusResponse(wiimStatus)

	jsonData, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal status")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
