package main

import (
	"fmt"
	"net/http"
)

// A middleware used for panic recovering when a server gets a 500 Internal Server Error
func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.errorResponse(w, r, http.StatusInternalServerError, fmt.Errorf("%s", err))

			}

		}()
		next.ServeHTTP(w, r)
	})
}
