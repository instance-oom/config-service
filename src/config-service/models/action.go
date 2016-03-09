package models

//Action : msg handler request body
type Action struct {
	ActionName  string `json:"actionName"`
	SystemName  string `json:"systemName"`
	ConfigName  string `json:"configName"`
	ConfigValue string `json:"configValue"`
}
