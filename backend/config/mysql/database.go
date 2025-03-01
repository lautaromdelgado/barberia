package mysql

import (
	"gorm.io/gorm"
)

type IConfig interface {
	GetDNS() string
	GetDriver() string
	GetDB() *gorm.DB
}

type Config struct {
	Connection *gorm.DB
}

func (c *Config) GetDNS() string {
	return ""
}
