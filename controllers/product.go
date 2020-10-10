package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"expodb/models"
	"strconv"
)

type ProductListController struct {
	beego.Controller
}

type ProductDetailController struct {
	beego.Controller
}

func  ( c *ProductListController) Get(){
	o := orm.NewOrm()

	// 获取 QuerySeter 对象，user 为表名
	qs := o.QueryTable("product")

	products := make([]models.Product,0)
	qs.All(&products)
	resp := make([]models.ProductResp,0)
	for _,v := range products{
		var product models.ProductResp
		product.Id = v.Id
		product.Name = v.Name
		product.Description = v.Description
		product.Space =v.Space
		product.Pics = append(product.Pics, v.Pic1, v.Pic2, v.Pic3)
		resp = append(resp, product)
	}
	ret := models.ProductListResp{
		Code:0,
		Data:resp,
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}

func (c *ProductDetailController) Get(){
	id,_ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()

	product :=models.ProductDetail{Id:id}
	o.Read(&product)
	ret := models.ProductDetailResp{
		Code:0,
		Data: product,
	}
	c.Data["json"] = ret
	c.ServeJSON()
}