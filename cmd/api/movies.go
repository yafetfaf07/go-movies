package main

import (
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
		Genre   []string     `json:"genre"`
	}

	v := validator.New()

	v.Check(input.Title != "", "title", "must be provided")

	v.Check(len(input.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(input.Year != 0, "year", "must be provided")

	v.Check(input.Year >= 1888, "year", "must be greater than 1888")
	
	v.Check(input.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(input.Runtime != 0, "runtime", "must be provided")

	v.Check(input.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(input.Genre != nil, "genres", "must be provided")

	v.Check(len(input.Genre) >= 1, "genres", "must contain at least 1 genre")

	v.Check(len(input.Genre) <= 5, "genres", "must not contain more than 5 genres")

	v.Check(validator.Unique(input.Genre), "genres", "must not contain duplicate values")
	
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err := app.readJSON(w, r, &input)
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
