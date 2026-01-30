package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	js := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	marshalled, err := json.Marshal(js)

	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Unknown error occured", http.StatusInternalServerError)

	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshalled)
}
