package models

import (
    _ "github.com/mattn/go-sqlite3"
    // _ "github.com/go-sql-driver/mysql"
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
    // engine, err = xorm.NewEngine("mysql", "root:123456@/go_demo?charset=utf8")
    if err != nil {
        panic(err)
    }
    err = engine.Sync(
		new(Account),
        new(Acra),
	)

	if err != nil {
		panic(err)
	}
}
