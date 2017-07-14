package controllers

import (
	"github.com/revel/revel"
    "github.com/adamzfc/revel_demo/app/models"
)

type Account struct {
	*revel.Controller
}

func (a Account) Index(account *models.Account) revel.Result {
    account_list :=  account.SelectAll()
	return a.Render(account_list)
}
