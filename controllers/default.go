package controllers

import (
	"tododat/models"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	u := models.User{Id: 1}
	tasks := models.ReadTasksFromUser(u)
	// for i := 0; i < len(list); i++ {
	// 	fmt.Println("'", list[i].Id, "' '", list[i].Content, "'", list[i].ProjectId, "'", list[i].CreatedAt, "'", list[i].CreatedBy)
	// }
	c.Data["tasks"] = tasks
	c.TplName = "index.tpl"
}
