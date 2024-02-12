package model

import (
	"github.com/astaxie/beego/orm"
)

type LsConfig struct {
	Id         int    `orm:"column(id);auto;pk"`
	ConfigType string `orm:"column(configType);size(100)" description:"config type"`
	Key        string `orm:"column(key);size(100)" description:"config name"`
	Value      string `orm:"column(value);size(256)" description:"config value"`
}

// func (config *LsConfig) TableName() string {
// 	return "ls_config"
// }
// GetConfig ... get the config by its key
// returns - config map
// error- err
func GetConfigByType(configType string) (map[string]string, error) {
	configs := []LsConfig{}
	o := orm.NewOrm()
	_, err := o.QueryTable(new(LsConfig)).Filter("configType", configType).All(&configs)
	if err != nil {
		return nil, err
	}
	config := make(map[string]string)
	for _, val := range configs {
		config[val.Key] = val.Value
	}
	return config, nil
}
