package handlers

import (
	"encoding/json"
	"net/http"
)

func TrackInceptionHandler(w http.ResponseWriter, r *http.Request, soundtrackUrls []string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(soundtrackUrls)
}
