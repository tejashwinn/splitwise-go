package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tejashwinn/splitwise/constants"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
}
