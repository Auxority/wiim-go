package api

type StatusResponse struct {
	Album             string `json:"album"`
	Artist            string `json:"artist"`
	Channel           string `json:"channel"`
	PlaybackMode      string `json:"playbackMode"`
	RepeatMode        string `json:"repeatMode"`
	Status            string `json:"status"`
	Title             string `json:"title"`
	Vendor            string `json:"vendor"`
	CurrentTrackIndex int    `json:"currentTrackIndex"`
	EqPreset          int    `json:"eqPreset"`
	TotalTracks       int    `json:"totalTracks"`
	TrackDurationMs   int    `json:"trackDurationMs"`
	TrackProgressMs   int    `json:"trackProgressMs"`
	Volume            int    `json:"volume"`
	IsMuted           bool   `json:"isMuted"`
}
