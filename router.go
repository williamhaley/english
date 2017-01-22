package main

import (
	"github.com/gorilla/mux"
	"github.com/williamhaley/negroni"

	"./controller"
	"./middleware"

	"fmt"
	"net/http"
)

func NewRouter() *negroni.Negroni {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/", Index).Methods("GET")

	wordController := controller.NewWordController()

	// TODO WFH Why does this only work on mux.NewRouter()? Why can't it
	// be a subrouter of `router`? That feels weird. Two routers. Why?
	apiRouter := mux.NewRouter().PathPrefix("/api").Subrouter()

	apiRouterV1 := apiRouter.PathPrefix("/v1").Subrouter()
	apiRouterV1.HandleFunc("/words", wordController.Create).Methods("POST")
	apiRouterV1.HandleFunc("/words", wordController.Index).Methods("GET")
	apiRouterV1.HandleFunc("/words/{wordId}", wordController.Show).Methods("GET")

	apiRouterV2 := apiRouter.PathPrefix("/v2").Subrouter()
	apiRouterV2.HandleFunc("/words", wordController.Create).Methods("POST")
	apiRouterV2.HandleFunc("/words", wordController.Index).Methods("GET")
	apiRouterV2.HandleFunc("/words/{wordId}", wordController.Show).Methods("GET")

	apiRouterV3 := apiRouter.PathPrefix("/v3").Subrouter()
	apiRouterV3.HandleFunc("/words", wordController.Create).Methods("POST")
	apiRouterV3.HandleFunc("/words", wordController.Index).Methods("GET")
	apiRouterV3.HandleFunc("/words/{wordId}", wordController.Show).Methods("GET")

	router.PathPrefix("/api/v1").Handler(negroni.New(
		middleware.NewAPIMiddleware("common api middleware"),
		middleware.NewAPIMiddleware("common auth middleware"),
		middleware.NewAPIMiddleware("specific to v1"),
		negroni.Wrap(apiRouter),
	))

	router.PathPrefix("/api/v2").Handler(negroni.New(
		middleware.NewAPIMiddleware("common api middleware"),
		middleware.NewAPIMiddleware("common auth middleware"),
		middleware.NewAPIMiddleware("shared with v1 & v2"),
		middleware.NewAPIMiddleware("specific to v2"),
		middleware.NewAPIMiddleware("common output middleware"),
		negroni.Wrap(apiRouter),
	))

	router.PathPrefix("/api/v3").Handler(negroni.New(
		middleware.NewAPIMiddleware("common api middleware"),
		middleware.NewAPIMiddleware("common auth middleware"),
		middleware.NewAPIMiddleware("shim for some reason"),
		middleware.NewAPIMiddleware("shared with v1 & v2"),
		middleware.NewAPIMiddleware("specific to v3"),
		middleware.NewAPIMiddleware("common output middleware"),
		negroni.Wrap(apiRouter),
	))

	// All of this middleware runs before any of the per-route middleware. Note
	// that it may look confusing, because the logger prints some info, then
	// passes to the next middleware, and after *all* middleware has run, prints
	// another timestamp and bit of info.
	negroniRunner := negroni.New()
	negroniRunner.Use(middleware.NewBaseMiddleware1("base 1"))
	negroniRunner.Use(middleware.NewLoggerMiddleware())
	negroniRunner.Use(middleware.NewBaseMiddleware1("base 2"))
	negroniRunner.UseHandler(router)

	return negroniRunner
}

func Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Welcome!")
}
