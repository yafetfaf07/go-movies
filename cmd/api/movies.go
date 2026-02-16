package main

import (
	"fmt"
	"net/http"
	"time"

	data "github.com/yafetfaf07/go-movies/internal/data"
	validator "github.com/yafetfaf07/go-movies/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genre"`
	}

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}
	v := validator.New()
	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)

	err:= app.models.Movies.Insert(movie)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		// app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		app.badRequestResponse(w, r, err)
		return
	}
	// fmt.Sprintf("%+v\n", input)
	err = app.writeJSON(w, http.StatusCreated, input, nil)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *application) getMovieByIDHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.GetIdFromParams(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}
	movies := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movies}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Server error try again later", http.StatusInternalServerError)
	}

}
