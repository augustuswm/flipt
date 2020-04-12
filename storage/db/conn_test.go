package db

import (
	"io/ioutil"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestEscapesForMySQL(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)
	db, _, err := Open("file:../../flipt_test.db")

	if err != nil {
		logger.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			logger.Fatal(err)
		}
	}()

	conn := NewConn(sq.StatementBuilder, db, MySQL)

	assert.Equal(t, "`key`", conn.Esc("key"))
}

func TestDoesNotEscapeForPostgres(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)
	db, _, err := Open("file:../../flipt_test.db")

	if err != nil {
		logger.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			logger.Fatal(err)
		}
	}()

	conn := NewConn(sq.StatementBuilder, db, Postgres)

	assert.Equal(t, "\"key\"", conn.Esc("key"))
}

func TestDoesNotEscapeForSQLite(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)
	db, _, err := Open("file:../../flipt_test.db")

	if err != nil {
		logger.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			logger.Fatal(err)
		}
	}()

	conn := NewConn(sq.StatementBuilder, db, SQLite)

	assert.Equal(t, "\"key\"", conn.Esc("key"))
}
