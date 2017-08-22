package model

type Word struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
	Type       string `json:"type"`
}
