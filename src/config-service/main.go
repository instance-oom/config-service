package main

import (
	"config-service/controllers"
	"config-service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	models.Init()

	beego.ErrorController(&controllers.ErrorController{})
	//Route Setting
	beego.Router("/config/v1/message-handler", &controllers.MsgController{})
	beego.Router("/config/v1/:system", &controllers.SystemController{})
	beego.Router("/config/v1/:system/:config", &controllers.ConfigController{})

	beego.Run()
}
