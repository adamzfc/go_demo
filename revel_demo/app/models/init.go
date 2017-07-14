package models

import (
    _ "github.com/mattn/go-sqlite3"
    "github.com/go-xorm/xorm"
    "github.com/revel/revel"
)

var engine *xorm.Engine

func init() {
    revel.OnAppStart(InitDB)
}

func InitDB() {
    var err error
    engine, err = xorm.NewEngine("sqlite3", "./test.db")
    if err != nil {
        panic(err)
    }
    err = engine.Sync(
		new(Account),
	)

	if err != nil {
		panic(err)
	}
}
