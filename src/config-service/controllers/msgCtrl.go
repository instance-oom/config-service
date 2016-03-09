package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

	"config-service/models"
)

//MsgController : handler msg request
type MsgController struct {
	BaseController
}

var action models.Action

//Prepare : validate request body
func (m *MsgController) Prepare() {
	if m.Ctx.Request.Method == "PUT" {
		if string(m.Ctx.Input.RequestBody) == "" {
			m.abort(400, "Request body cannot be null or empty.")
		}
		err := json.Unmarshal(m.Ctx.Input.RequestBody, &action)
		if err != nil {
			m.abort(400, fmt.Sprintf("Invalidate request json data. %s", err.Error()))
		}
		switch strings.ToLower(action.ActionName) {
		case "createsystem":
		case "deletesystem":
			if action.SystemName == "" {
				m.abort(400, "'systemName' cannot be null or empty.")
			}
			break
		case "createconfig":
		case "updateconfig":
		case "deleteconfig":
			msg := []string{}
			if action.SystemName == "" {
				msg = append(msg, "'systemName' cannot be null or empty.")
			}
			if action.ConfigName == "" {
				msg = append(msg, "'configName' cannot be null or empty.")
			}
			if len(msg) > 0 {
				m.abort(400, msg)
			}
			break
		default:
			m.abort(400, "'actionName' incorrect.['CreateSystem'、'DeleteSystem'、'CreateConfig'、'UpdateConfig'、'DeleteConfig']")
			break
		}
	}
}

//Put : handle put request
func (m *MsgController) Put() {
	var err error
	configNode := models.ConfigNode{
		SystemName:  action.SystemName,
		ConfigName:  action.ConfigName,
		ConfigValue: action.ConfigValue,
	}
	switch strings.ToLower(action.ActionName) {
	case "createsystem":
		err = models.CreateSystem(action.SystemName)
		break
	case "deletesystem":
		err = models.DeleteSystem(action.SystemName)
		break
	case "createconfig":
		err = models.CreateConfig(configNode)
		break
	case "updateconfig":
		err = models.UpdateConfig(configNode)
		break
	case "deleteconfig":
		err = models.DeleteConfig(action.SystemName, action.ConfigName)
		break
	}
	if err != nil {
		m.abort(500, err.Error())
	}
	m.Ctx.ResponseWriter.Header().Add("Content-Type", "application/json")
	m.Ctx.ResponseWriter.Write([]byte(""))
}
