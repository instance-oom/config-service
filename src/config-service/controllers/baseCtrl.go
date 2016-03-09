package controllers

import (
	"github.com/astaxie/beego"
)

//BaseController : provider some common func
type BaseController struct {
	beego.Controller
}

func (b *BaseController) abort(code int, msg interface{}) {
	b.Data["json"] = map[string]interface{}{
		"code":  code,
		"error": msg,
	}
	b.Ctx.Output.SetStatus(code)
	b.ServeJSON(false)
	b.StopRun()
}
