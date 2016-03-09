package models

import "fmt"

//ConfigNode : config node info
type ConfigNode struct {
	SystemName  string `json:"SystemName"`
	ConfigName  string `json:"ConfigName"`
	ConfigValue string `json:"ConfigValue"`
}

//GetConfig : get config node info
func GetConfig(system, config string) (*ConfigNode, error) {
	path := fmt.Sprintf("/%s/%s", system, config)
	value, _, err := zkClient.Get(path)
	if err != nil {
		return nil, err
	}
	configNode := &ConfigNode{
		SystemName:  system,
		ConfigName:  config,
		ConfigValue: string(value),
	}
	return configNode, nil
}

//CreateConfig : create config node
func CreateConfig(configNode ConfigNode) error {
	path := fmt.Sprintf("/%s/%s", configNode.SystemName, configNode.ConfigName)
	_, err := zkClient.Create(path, []byte(configNode.ConfigValue), flags, acl)
	if err != nil {
		return err
	}
	return nil
}

//UpdateConfig : update config node
func UpdateConfig(configNode ConfigNode) error {
	path := fmt.Sprintf("/%s/%s", configNode.SystemName, configNode.ConfigName)
	_, err := zkClient.Set(path, []byte(configNode.ConfigValue), -1)
	if err != nil {
		return err
	}
	return nil
}

//DeleteConfig : delete config node
func DeleteConfig(system, config string) error {
	path := fmt.Sprintf("/%s/%s", system, config)
	err := zkClient.Delete(path, -1)
	if err != nil {
		return err
	}
	return nil
}
