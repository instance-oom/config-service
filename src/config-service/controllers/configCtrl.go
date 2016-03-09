package controllers

import (
	"config-service/models"
)

//ConfigController : hangle config request
type ConfigController struct {
	BaseController
}

//Get : get config
func (c *ConfigController) Get() {
	system := c.Ctx.Input.Param(":system")
	config := c.Ctx.Input.Param(":config")
	configNode, err := models.GetConfig(system, config)
	if err != nil {
		c.abort(500, err.Error())
	}
	c.Data["json"] = configNode
	c.ServeJSON()
}
