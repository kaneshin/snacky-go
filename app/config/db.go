package config

import "bytes"

type DbConfig struct {
	User     string
	Password string
	Protocol string
	Address  string
	Port     string
	DbName   string
}

func NewDbConfig() *DbConfig {
	dbConfig := DbConfig{}
	dbConfig.User = GetString("db.user", "root")
	dbConfig.Password = GetString("db.password", "")
	dbConfig.Protocol = GetString("db.protocol", "tcp")
	dbConfig.Address = GetString("db.address", "localhost")
	dbConfig.Port = GetString("db.port", "")
	dbConfig.DbName = GetString("db.dbname", "")
	return &dbConfig
}

func (c *DbConfig) DSN() string {
	var buffer bytes.Buffer

	buffer.WriteString(c.User)
	if len(c.Password) != 0 {
		buffer.WriteString(":")
		buffer.WriteString(c.Password)
	}
	buffer.WriteString("@")
	buffer.WriteString(c.Protocol)
	buffer.WriteString("(")
	buffer.WriteString(c.Address)
	if len(c.Port) != 0 {
		buffer.WriteString(":")
		buffer.WriteString(c.Port)
	}
	buffer.WriteString(")/")
	buffer.WriteString(c.DbName)
	return buffer.String()
}
