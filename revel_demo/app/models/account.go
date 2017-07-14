package models

import (
	"github.com/revel/revel"
)

type Account struct {
    Id int64 `xorm:"pk autoincr"`
    Name string `xorm:"char(32)"`
}

func (a Account) SelectAll() (account_arr []Account) {
    account_list := []Account{}
    engine.Find(&account_list)
    return account_list
}

func (a Account) InsertOne(name string) bool {
    account := new(Account)
    account.Name = name
    has, err := engine.Table("account").Insert(account)
    if err != nil {
        revel.WARN.Println(has)
        revel.WARN.Printf("Error %v", err)
        return false
    }
    return true
}

func (a Account) SelectById(id string) Account {
    ac := new(Account)
    has, err := engine.Where("id=?", id).Get(ac)
    if err != nil || !has{
        (*ac).Id = -1
        return *ac
    } else {
        return *ac
    }
}
