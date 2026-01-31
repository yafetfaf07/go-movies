package main

import (
	"fmt"
	"net/http"
	"time"

	data "github.com/yafetfaf07/go-movies/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) getMovieByIDHandler(w http.ResponseWriter, r *http.Request) {

id,err:=app.GetIdFromParams(r)

if err!=nil {
	http.NotFound(w,r)
	return
}
movies:=data.Movie{
	ID: id,
	CreatedAt: time.Now(),
	Title: "Casablanca",
	Runtime: 102,
	Genres: []string{"drama", "romance","war"},
	Version: 1,
}
err=app.writeJSON(w,http.StatusOK,envelope{"movie":movies},nil)
if err!=nil {
	app.logger.Error(err.Error())
	http.Error(w,"Server error try again later",http.StatusInternalServerError)
}


}
