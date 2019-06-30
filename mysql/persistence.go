package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/janstun/gear"
)

func NewRelationalSqlDatabase(dsn string) (gear.RelationalSqlDatabase, error) {
	return sql.Open("mysql", dsn)
}
