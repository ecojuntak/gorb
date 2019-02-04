package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	logrus.Errorln(message)
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)

	if err != nil {
		logrus.Errorln(err)
	}
}
