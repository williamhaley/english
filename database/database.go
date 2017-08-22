package database

import (
	"os"

	"../model"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

func NewDatabase() (Database, error) {
	db, err := Initialize()
	if err != nil {
		return nil, err
	}

	return &database{
		db: db,
	}, nil
}

type Database interface {
	GetAllWords() ([]*model.Word, error)
	GetNonVerbs() ([]*model.Word, error)
}

type database struct {
	db *sqlx.DB
}

func (d *database) GetAllWords() ([]*model.Word, error) {
	var words []*model.Word

	err := d.db.Select(&words, `SELECT term, definition, type FROM words`)
	if err != nil {
		log.WithError(err).Error("Error getting all words")
		return nil, err
	}

	return words, nil
}

func (d *database) GetNonVerbs() ([]*model.Word, error) {
	var words []*model.Word

	// This query is purposefully ugly. Testing out `IN`.
	err := d.db.Select(&words, `
		SELECT 
			term, definition, type
		FROM
			words
		WHERE
			type IN (?, ?)
	`, "noun", "adjective")
	if err != nil {
		log.WithError(err).Error("Error getting all words")
		return nil, err
	}

	return words, nil
}

var schema = `
CREATE TABLE words (
    term VARCHAR(512),
    definition TEXT,
    type VARCHAR(128)
);
INSERT INTO words (
	term, definition, type
) VALUES (
	"dog", "a canine", "noun"
), (
	"cat", "a feline", "noun"
), (
	"run", "move quickly", "verb"
), (
	"fly", "soar in the sky", "verb"
), (
	"big", "to be large", "adjective"
), (
	"small", "to be tiny", "adjective"
), (
	"house", "a swelling", "noun"
), (
	"boat", "an aquatic vehicle", "noun"
);
`

func Initialize() (*sqlx.DB, error) {
	var isFirstRun bool = false

	_, err := os.Stat("english.db")
	if os.IsNotExist(err) {
		isFirstRun = true
	} else if err != nil {
		log.Panic("Error while attempting to open database", err)
		return nil, err
	}

	db, err := sqlx.Open("sqlite3", "./english.db")
	if err != nil {
		log.Panic("Error opening database", err)
		return nil, err
	}

	if isFirstRun {
		log.Info("First run. Creating schema")
		_, err := db.Exec(schema)
		if err != nil {
			log.Panic("Error creating schema")
		}
	}

	return db, nil
}
