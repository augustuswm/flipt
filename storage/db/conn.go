package db

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/markphelps/flipt/storage"
)

var _ storage.Conn = &Conn{}

type Conn struct {
	builder sq.StatementBuilderType
	db      *sql.DB
	driver  Driver
}

func NewConn(builder sq.StatementBuilderType, db *sql.DB, driver Driver) *Conn {
	return &Conn{
		builder: builder,
		db:      db,
		driver:  driver,
	}
}

func (c *Conn) Esc(term string) string {
	if c.driver == MySQL {
		return fmt.Sprintf("`%v`", term)
	} else {
		return term
	}
}
