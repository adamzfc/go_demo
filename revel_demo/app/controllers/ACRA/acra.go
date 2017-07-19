package controllers

import (
	"github.com/revel/revel"
    "github.com/adamzfc/revel_demo/app/models"
)

type Acra struct {
	*revel.Controller
}

func (c Acra) Index() revel.Result {
	return c.RenderJSON("{}")
}

func (c Acra) AddReport(ac *models.Acra) revel.Result {
    var acra models.Acra;
    c.Params.BindJSON(&acra)
    ac.InsertOne(acra)
    return c.RenderJSON("{}")
}
