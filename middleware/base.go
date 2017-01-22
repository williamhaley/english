package middleware

import (
	"fmt"
	"net/http"
)

type baseMiddleware1 struct {
	name string
}

func (m *baseMiddleware1) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Base Middleware:", m.name)
	next(rw, r)
}

func NewBaseMiddleware1(name string) *baseMiddleware1 {
	return &baseMiddleware1{
		name: name,
	}
}
