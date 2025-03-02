package variables_entorno

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	User      string
	Password  string
	Host      string
	Port      int
	DBName    string
	Charset   string
	ParseTime bool
	Loc       string
}

// Variable global para almacenar la configuración (Singleton)
var (
	configInstance *Config
	once           sync.Once
)

// GetConfig obtiene la configuración de la base de datos
func (c *Config) GetConfig() (*Config, error) {
	var err error
	once.Do(func() {
		err = godotenv.Load("../config/variables_entorno/.env")
		if err != nil {
			err = errors.New("error loading .env file")
		}

		port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
		parseTime, _ := strconv.ParseBool(os.Getenv("DB_PARSE_TIME"))

		configInstance = &Config{
			User:      os.Getenv("DB_USER"),
			Password:  os.Getenv("DB_PASSWORD"),
			Host:      os.Getenv("DB_HOST"),
			Port:      port,
			DBName:    os.Getenv("DB_NAME"),
			Charset:   os.Getenv("DB_CHARSET"),
			ParseTime: parseTime,
			Loc:       os.Getenv("DB_LOC"),
		}
	})

	if configInstance == nil {
		return nil, errors.New("error loading configuration")
	}

	return configInstance, nil
}

// GetDNS obtiene el DSN de la base de datos
func (c Config) GetDNS() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		url.QueryEscape(c.User),
		url.QueryEscape(c.Password),
		c.Host,
		c.Port,
		c.DBName,
		c.Charset,
		c.ParseTime,
		c.Loc,
	)

	return dsn
}
