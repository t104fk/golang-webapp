package db

import (
	_ "github.com/lib/pq" // for sql.Open
	"github.com/go-gorp/gorp"
	"database/sql"
	"golang-webapp/model"
	"fmt"
	"golang-webapp/config"
)

var Cfg *Database

type Database struct {
	DbMap *gorp.DbMap
}

func InitDb() {
	Cfg = &Database{}

	datasource := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
					config.Cfg.DbHost, config.Cfg.DbUser, config.Cfg.DbName, config.Cfg.DbPass)
	db, err := sql.Open("postgres", datasource)
	if err != nil {
		// handle error
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(model.Article{}, "articles").SetKeys(true, "Id")
	dbmap.AddTableWithName(model.User{}, "users").SetKeys(true, "Id")
	Cfg.DbMap = dbmap
}
