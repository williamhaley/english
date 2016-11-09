package model

import (
	"errors"
)

type Word interface{}
type word struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
}

func NewWord(properties map[string]interface{}) (Word, error) {
	word := &word{}

	var ok bool

	if word.Term, ok = properties["term"].(string); !ok {
		return nil, errors.New("Invalid value for `term`.")
	}

	if word.Definition, ok = properties["definition"].(string); !ok {
		return nil, errors.New("Invalid value for `definition`.")
	}

	return word, nil
}
