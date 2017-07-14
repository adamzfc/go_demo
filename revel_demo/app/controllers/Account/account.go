package controllers

import (
	"github.com/revel/revel"
    "github.com/adamzfc/revel_demo/app/models"
)

type Account struct {
	*revel.Controller
}

func (c Account) Index(account *models.Account) revel.Result {
    account_list :=  account.SelectAll()
	return c.Render(account_list)
}

func (c Account) Add(account *models.Account, accountName string) revel.Result {
    c.Validation.Required(accountName).Message("Account name is required!")
    c.Validation.MinSize(accountName, 3).Message("Account name is not long enough!")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(Account.Index)
    }
    if account.InsertOne(accountName) {
        c.Flash.Success("Add Success!")
    } else {
        c.Flash.Error("Add Failed!")
    }
    return c.Redirect(Account.Index)
}

func (c Account) ApiList(account *models.Account) revel.Result {
    id := c.Params.Query.Get("id")
    if len(id) > 0 {
        a := account.SelectById(id)
        if a.Id != -1 {
            return c.RenderJSON(a)
        } else {
            return c.RenderJSON(a)
        }
    }
    account_list :=  account.SelectAll()
	return c.RenderJSON(account_list)
}

