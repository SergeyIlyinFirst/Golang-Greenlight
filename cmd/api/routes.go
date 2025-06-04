package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /v1/healthcheck", app.recoverPanic(http.HandlerFunc(app.healthcheckHandler)))
	router.Handle("GET /v1/movies", app.recoverPanic(http.HandlerFunc(app.listMoviesHandler)))
	router.Handle("POST /v1/movies", app.recoverPanic(http.HandlerFunc(app.createMovieHandler)))
	router.Handle("GET /v1/movies/{id}", app.recoverPanic(http.HandlerFunc(app.showMovieHandler)))
	router.Handle("PATCH /v1/movies/{id}", app.recoverPanic(http.HandlerFunc(app.updateMovieHandler)))
	router.Handle("DELETE /v1/movies/{id}", app.recoverPanic(http.HandlerFunc(app.deleteMovieHandler)))

	return router
}
