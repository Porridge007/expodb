package main

import (
	_ "expodb/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"expodb/models"
)

func init(){
	orm.RegisterDataBase("default", "mysql", "root:123456!@tcp(127.0.0.1:3306)/expo?charset=utf8&loc=Asia%2FShanghai", 30)
	orm.RegisterModel(new(models.Product))
	orm.RegisterModel(new(models.ProductDetail))
	orm.RegisterModel(new(models.Expo))
	orm.RegisterModel(new(models.ExpoDetail))

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
