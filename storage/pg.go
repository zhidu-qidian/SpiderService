package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"workspace/SpiderService/config"
)

var PG *sqlx.DB

func init() {
	db, err := sqlx.Open("postgres", config.C.Postgres.String())
	db.SetMaxOpenConns(config.C.Postgres.MaxConnections)
	db.SetMaxIdleConns(config.C.Postgres.MaxIdles)
	if err != nil {
		panic(err)
	}
	PG = db
}
