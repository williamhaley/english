package controller

import (
	"encoding/json"
	"net/http"
)

func Write(rw http.ResponseWriter, status int, response interface{}) {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(response)
}

func WriteError(rw http.ResponseWriter, status int, err error) {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(status)

	type errorStruct struct {
		Error string `json:"error"`
	}

	json.NewEncoder(rw).Encode(errorStruct{
		Error: err.Error(),
	})
}
