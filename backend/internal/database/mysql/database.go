package mysql

import (
	"barberia/config/variables_entorno"
	"errors"
	"fmt"
	"log"

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
		log.Printf("Error getting DSN: %v", err)
		return errors.New("error getting dsn")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Detailed connection error: %v", err)
		return fmt.Errorf("error opening connection: %v", err)
	}
	c.Connection = db
	return nil
}

func (c *Config) GetDB() (*gorm.DB, error) {
	if err := c.OpenConnection(); err != nil {
		log.Printf("GetDB error: %v", err)
		return nil, errors.New("error opening connection: GetDB")
	}
	if c.Connection == nil {
		return nil, errors.New("error getting connection")
	}
	return c.Connection, nil
}
