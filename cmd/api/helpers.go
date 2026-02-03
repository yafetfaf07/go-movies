package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]any

func (app *application) GetIdFromParams(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("Id is invalid")
	}
	return id, nil

}

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// Encode the data to JSON, returning the error if there was one.
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// Append a newline to make it easier to view in terminal applications.
	js = append(js, '\n')
	// At this point, we know that we won't encounter any more errors before writing the
	// response, so it's safe to add any headers that we want to include. We loop
	// through the header map and add each header to the http.ResponseWriter header map.
	// Note that it's OK if the provided header map is nil. Go doesn't throw an error
	// if you try to range over (or generally, read from) a nil map.
	for key, value := range headers {
		w.Header()[key] = value
	}
	// Add the "Content-Type: application/json" header, then write the status code and
	// JSON response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func(app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
err:=json.NewDecoder(r.Body).Decode(dst)

if err!=nil {
	var syntaxError *json.SyntaxError
	// var unmarshalTypeError *json.UnmarshalTypeError
	// var invalidUnmarshalError *json.InvalidUnmarshalError

	switch {
	case errors.As(err,&syntaxError) :
		return errors.New("body contains badly formed JSON")
	}

}
return nil
}