package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gh0sty02/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParams(r)

	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Duration:  102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envolope{"movie": movie}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	// fmt.Fprintf(w, "show the details of movie : %d\n", id)
}
