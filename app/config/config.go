package config

import "github.com/revel/revel"

type Default struct {
	Db *DbConfig
}

var appDefault *Default

func GetDefault() *Default {
	if appDefault != nil {
		return appDefault
	}
	appDefault = &Default{}
	appDefault.Db = NewDbConfig()
	return appDefault
}

func GetString(key string, def string) string {
	if val, ok := revel.Config.String(key); ok {
		return val
	}
	return def
}

func GetInt(key string, def int) int {
	if val, ok := revel.Config.Int(key); ok {
		return val
	}
	return def
}
