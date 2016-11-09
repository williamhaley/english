package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"./controller"

	"fmt"
	"net/http"
	"time"
)

func NewRouter() *negroni.Negroni {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/", Index).Methods("GET")

	wordController := controller.NewWordController()

	// TODO WFH Why does this only work on mux.NewRouter()? Why can't it
	// be a subrouter of `router`? That feels weird. Two routers. Why?
	apiRouter := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/words", wordController.Create).Methods("POST")
	apiRouter.HandleFunc("/words", wordController.Index).Methods("GET")
	apiRouter.HandleFunc("/words/{wordId}", wordController.Show).Methods("GET")

	router.PathPrefix("/api/v1").Handler(negroni.New(
		&middleware1{},
		negroni.Wrap(apiRouter),
	))

	negroniRunner := negroni.New()
	negroniRunner.Use(&baseMiddleware1{})
	negroniRunner.Use(&loggerMiddleware{})
	negroniRunner.UseHandler(router)

	return negroniRunner
}

type middleware1 struct{}

func (m *middleware1) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Middleware1")
	next(rw, r)
}

type baseMiddleware1 struct{}

func (m *baseMiddleware1) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Base Middleware 1")
	next(rw, r)
}

type loggerMiddleware struct{}

func (m *loggerMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	format := "2006-01-02 15:04:05"

	start := time.Now()
	fmt.Printf("> %v | %v %v\n", start.UTC().Format(format), r.Method, r.URL.Path)

	next(rw, r)

	end := time.Now()
	fmt.Printf("< %v | %v\n", end.UTC().Format(format), end.Sub(start))
}

func Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Welcome!")
}
