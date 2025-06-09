package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /v1/healthcheck", app.recoverPanic(app.rateLimit(app.authenticate(http.HandlerFunc(app.healthcheckHandler)))))

	router.Handle("GET /v1/movies", app.recoverPanic(app.rateLimit(app.authenticate(app.requireActivatedUser(http.HandlerFunc(app.listMoviesHandler))))))
	router.Handle("POST /v1/movies", app.recoverPanic(app.rateLimit(app.authenticate(app.requireActivatedUser(http.HandlerFunc(app.createMovieHandler))))))
	router.Handle("GET /v1/movies/{id}", app.recoverPanic(app.rateLimit(app.authenticate(app.requireActivatedUser(http.HandlerFunc(app.showMovieHandler))))))
	router.Handle("PATCH /v1/movies/{id}", app.recoverPanic(app.rateLimit(app.authenticate(app.requireActivatedUser(http.HandlerFunc(app.updateMovieHandler))))))
	router.Handle("DELETE /v1/movies/{id}", app.recoverPanic(app.rateLimit(app.authenticate(app.requireActivatedUser(http.HandlerFunc(app.deleteMovieHandler))))))

	router.Handle("POST /v1/users", app.recoverPanic(app.rateLimit(app.authenticate(http.HandlerFunc(app.registerUserHandler)))))
	router.Handle("PUT /v1/users/activated", app.recoverPanic(app.rateLimit(app.authenticate(http.HandlerFunc(app.activateUserHandler)))))

	router.Handle("POST /v1/tokens/authentication", app.recoverPanic(app.rateLimit(app.authenticate(http.HandlerFunc(app.createAuthenticationTokenHandler)))))

	return router
}
