package mysql

import (
	"barberia/config/variables_entorno"
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Connection *gorm.DB
}

func (c *Config) OpenConnection() error {
	var configEnv = new(variables_entorno.Config)
	dsn, err := configEnv.GetDNS()
	if err != nil {
		return errors.New("error getting dsn")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("error opening connection")
	}
	c.Connection = db
	return nil
}

func (c *Config) GetDB() (*gorm.DB, error) {
	if err := c.OpenConnection(); err != nil {
		return nil, errors.New("error opening connection: GetDB")
	}
	if c.Connection == nil {
		return nil, errors.New("error getting connection")
	}
	return c.Connection, nil
}
