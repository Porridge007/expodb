package controllers

import (
	"github.com/astaxie/beego"
	"expodb/models"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type ExpoListController struct {
	beego.Controller
}

type ExpoDetailController struct {
	beego.Controller
}

func (c *ExpoListController) Get(){
	o := orm.NewOrm()
	qs := o.QueryTable("expo")

	expo := make([]models.Expo,0)
	qs.All(&expo)
	ret := models.ExpoListResp{
		Code:0,
		Data:expo,
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}

func (c *ExpoDetailController) Get(){
	id,_ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()

	expo :=models.ExpoDetail{Id:id}
	o.Read(&expo)
	ret := models.ExpoDetailResp{
		Code:0,
		Data: expo,
	}
	c.Data["json"] = ret
	c.ServeJSON()
}