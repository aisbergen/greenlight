package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/shynggys9219/greenlight/internal/data"
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint.
// return a JSON response.
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	//Declare an anonymous struct to hold the information that we expect to be in the
	// HTTP request body (note that the field names and types in the struct are a subset
	// of the Movie struct that we created earlier). This struct will be our *target
	// decode destination*.
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	err := app.readJSON(w, r, &input) 
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	err = app.models.Movies.Insert(movie)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"movie": movie}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	movie,err:=app.models.Movies.Get(id)
	if err !=nil{
		switch{
		case errors.Is(err,data.ErrRecordNotFound):
			app.notFoundResponse(w,r)
		default:
			app.serverErrorResponse(w,r,err)
		}
		return
	}
	err=app.writeJSON(w,http.StatusOK, envelope{"movie":movie},nil)
	if err!=nil{
		app.serverErrorResponse(w,r,err)
	}

movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	

	func (app *application) deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
		id, err := app.readIDParam(r)
		if err != nil {
		app.notFoundResponse(w, r)
		return
		}
		err = app.models.Movies.Delete(id)
		if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
		app.notFoundResponse(w, r)
		default:
		app.serverErrorResponse(w, r, err)
		}
		return
		}
		err = app.writeJSON(w, http.StatusOK, envelope{"message": "movie successfully deleted"}, nil)
		if err != nil {
		app.serverErrorResponse(w, r, err)
		}
		}
		