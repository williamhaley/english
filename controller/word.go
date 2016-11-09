package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"../model"

	"github.com/gorilla/mux"
)

var words []model.Word = []model.Word{}

type WordController interface {
	Show(rw http.ResponseWriter, r *http.Request)
	Create(rw http.ResponseWriter, r *http.Request)
	Index(rw http.ResponseWriter, r *http.Request)
}
type wordController struct{}

func NewWordController() WordController {
	return &wordController{}
}

func (c *wordController) Create(rw http.ResponseWriter, r *http.Request) {
	var properties map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&properties); err != nil {
		Write(rw, http.StatusBadRequest, err)
		return
	}

	word, err := model.NewWord(properties)
	if err != nil {
		Write(rw, http.StatusBadRequest, err)
		return
	}

	words = append(words, word)

	Write(rw, http.StatusOK, word)
}

func (c *wordController) Index(rw http.ResponseWriter, r *http.Request) {
	Write(rw, http.StatusOK, words)
}

func (c *wordController) Show(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wordId, err := strconv.ParseInt(vars["wordId"], 10, 64)
	if err != nil {
		Write(rw, http.StatusBadRequest, err)
		return
	}

	if wordId < 0 || wordId >= int64(len(words)) {
		WriteError(rw, http.StatusNotFound, errors.New("Word not found"))
		return
	}

	Write(rw, http.StatusOK, words[wordId])
}
