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
var once sync.Once

// GetConfig obtiene la configuración de la base de datos
func (c *Config) GetConfig() error {
	var err error
	once.Do(func() {
		err = godotenv.Load("../config/variables_entorno/.env")
		if err != nil {
			err = errors.New("error loading .env file")
		}

		port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
		parseTime, _ := strconv.ParseBool(os.Getenv("DB_PARSE_TIME"))

		c.User = os.Getenv("DB_USER")
		c.Password = os.Getenv("DB_PASSWORD")
		c.Host = os.Getenv("DB_HOST")
		c.Port = port
		c.DBName = os.Getenv("DB_NAME")
		c.Charset = os.Getenv("DB_CHARSET")
		c.ParseTime = parseTime
		c.Loc = os.Getenv("DB_LOC")
	})

	if c == nil {
		return errors.New("error getting config: GetConfig is nil")
	}

	return nil
}

// GetDNS obtiene el DSN de la base de datos
func (c Config) GetDNS() (string, error) {
	if err := c.GetConfig(); err != nil {
		return "", err
	}
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

	return dsn, nil
}
