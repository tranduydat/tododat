package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"os"
	"strconv"
	_ "tododat/routers"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		beego.BConfig.Listen.HTTPPort = port
	}
	beego.Run()
}
