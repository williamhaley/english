package middleware

import (
	"fmt"
	"net/http"
)

type apiMiddleware struct {
	name string
}

func (m *apiMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("API Middleware:", m.name)
	next(rw, r)
}

func NewAPIMiddleware(name string) *apiMiddleware {
	return &apiMiddleware{
		name: name,
	}
}
