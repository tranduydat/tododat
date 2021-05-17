package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"tododat/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/done", &controllers.TaskController{Action: "done", Method: "post"})
	beego.Router("/add", &controllers.TaskController{Action: "add", Method: "post"})
	beego.SetStaticPath("/static", "static")
}
