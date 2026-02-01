package main

import "net/http"

// this is for error logging which will help a lot in debugging
func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"message": message}
	err := app.writeJSON(w, status, env, nil)

	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}

}

// This is for unknown crashes, database connection, third party crashes or internal problems
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "An unexpected error happen"

	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// for resources that are not found
func(app *application) dataNotFound(w http.ResponseWriter) {
	message:="The requested resource not found"
	app.writeJSON(w,http.StatusNotFound,message,nil)
}