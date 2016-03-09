package controllers

import (
	"config-service/models"
)

//SystemController : handle system request
type SystemController struct {
	BaseController
}

//Get : get config list
func (s *SystemController) Get() {
	system := s.Ctx.Input.Param(":system")
	systemNode, err := models.GetSystem(system)
	if err != nil {
		s.abort(500, err.Error())
	}
	s.Data["json"] = systemNode
	s.ServeJSON()
}
