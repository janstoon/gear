package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gitlab.com/janstun/gear"
)

func NewSql(dsn string) (gear.Sql, error) {
	if db, err := sql.Open("postgres", dsn); err != nil {
		return nil, fmt.Errorf("unable to create the database resource: %s", err)
	} else if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping the database: %s", err)
	} else {
		return db, nil
	}
}
