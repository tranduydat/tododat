package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"time"
	"tododat/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	u := models.User{
		Username:  "kimanh",
		Password:  "kimanh",
		Email:     "kimanh@gmail.com",
		CreatedAt: time.Now(),
	}
	var p models.Project
	p.ID = 1
	t := models.Task{
		Content:   "Test 1",
		ProjectId: &p,
		CreatedAt: time.Now(),
		CreatedBy: &u,
	}
	models.AddTask(t, p, u)
}
