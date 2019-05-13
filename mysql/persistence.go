package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/janstun/actor"
)

func NewRelationalSqlDatabase(dsn string) (actor.RelationalSqlDatabase, error) {
	return sql.Open("mysql", dsn)
}
