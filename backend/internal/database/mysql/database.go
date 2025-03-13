package mysql

import (
	"barberia/config/variables_entorno"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	// Configuración de logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Salida en consola
		logger.Config{
			SlowThreshold:             time.Second, // Marca consultas lentas (>1s)
			LogLevel:                  logger.Info, // Nivel de logs (Info, Warn, Error)
			IgnoreRecordNotFoundError: true,        // Ignorar "record not found"
			Colorful:                  true,        // Logs en color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Printf("Detailed connection error: %v", err)
		return fmt.Errorf("error opening connection: %v", err)
	}

	// Debug mode
	db = db.Debug()

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
