// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"expodb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("api/goods/list", &controllers.ProductListController{})
	beego.Router("api/goods/:id", &controllers.ProductDetailController{})

	beego.Router("api/expo/list", &controllers.ExpoListController{})
	beego.Router("api/expo/:id", &controllers.ExpoDetailController{})
}
