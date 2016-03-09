package models

import "fmt"

//SystemNode : system node info
type SystemNode struct {
	SystemName string   `json:"SystemName"`
	Configs    []string `json:"Configs"`
}

//CreateSystem : create system node
func CreateSystem(system string) error {
	path := fmt.Sprintf("/%s", system)
	_, err := zkClient.Create(path, []byte(""), flags, acl)
	if err != nil {
		return err
	}
	return nil
}

//GetSystem : get system node info with config list by name
func GetSystem(system string) (*SystemNode, error) {
	path := fmt.Sprintf("/%s", system)
	configs, _, err := zkClient.Children(path)
	if err != nil {
		return nil, err
	}
	systemNode := &SystemNode{
		SystemName: system,
		Configs:    configs,
	}
	return systemNode, nil
}

//DeleteSystem : delete system by name
func DeleteSystem(system string) error {
	systemNode, err := GetSystem(system)
	if err != nil {
		return err
	}
	for _, config := range systemNode.Configs {
		err := DeleteConfig(system, config)
		if err != nil {
			return err
		}
	}
	err = zkClient.Delete(fmt.Sprintf("/%s", system), -1)
	return err
}
