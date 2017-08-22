package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"../database"
	"../model"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var words []model.Word = []model.Word{}

type WordController interface {
	Show(rw http.ResponseWriter, r *http.Request)
	Create(rw http.ResponseWriter, r *http.Request)
	Index(rw http.ResponseWriter, r *http.Request)
	NonVerbs(rw http.ResponseWriter, r *http.Request)
}
type wordController struct {
	db database.Database
}

func NewWordController(db database.Database) WordController {
	return &wordController{db: db}
}

func (c *wordController) Create(rw http.ResponseWriter, r *http.Request) {
	var word model.Word
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&word); err != nil {
		log.WithError(err).Error("Error decoding parameters for new word")
		Write(rw, http.StatusBadRequest, err)
		return
	}

	words = append(words, word)

	Write(rw, http.StatusOK, word)
}

func (c *wordController) Index(rw http.ResponseWriter, r *http.Request) {
	words, err := c.db.GetAllWords()
	if err != nil {
		log.WithError(err).Error("Error retrieving words")
	}

	Write(rw, http.StatusOK, words)
}

func (c *wordController) NonVerbs(rw http.ResponseWriter, r *http.Request) {
	words, err := c.db.GetNonVerbs()
	if err != nil {
		log.WithError(err).Error("Error retrieving non-verbs")
	}

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
