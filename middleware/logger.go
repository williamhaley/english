package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type loggerMiddleware struct{}

func (m *loggerMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	format := "2006-01-02 15:04:05"

	start := time.Now()
	fmt.Printf("> %v | %v %v\n", start.UTC().Format(format), r.Method, r.URL.Path)

	next(rw, r)

	end := time.Now()
	fmt.Printf("< %v | %v\n", end.UTC().Format(format), end.Sub(start))
}

func NewLoggerMiddleware() *loggerMiddleware {
	return &loggerMiddleware{}
}
