package models

import (
	"github.com/revel/revel"
)

type Account struct {
    Id int64 `xorm:"pk"`
    Name string `xorm:"char(32)"`
}

func (a Account) SelectAll() (account_arr []Account) {
    account_list := []Account{}
    engine.Find(&account_list)
    return account_list
}

func (a *Account) InsertOne() bool {
    account := new(Account)
    account.Name = a.Name
    has, err := engine.Table("account").Insert(account)
    if err != nil {
        revel.WARN.Println(has)
        revel.WARN.Printf("Error %v", err)
        return false
    }
    return true
}

