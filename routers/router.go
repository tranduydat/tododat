package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"tododat/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.SetStaticPath("/static", "static")
}
