package libs

import (
	"fmt"
	"log"
	"os"

	"github.com/Unknowns24/backend-template/config"
	"github.com/Unknowns24/backend-template/internal/interfaces"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	Charset  string
}

var DB *gorm.DB

func StablishDatabaseConnection() {
	dbCfg := DbConfig{
		Host:     config.ENV.DB_HOST,
		Port:     config.ENV.DB_PORT,
		Database: config.ENV.DB_NAME,
		User:     config.ENV.DB_USER,
		Password: config.ENV.DB_PASSWORD,
		Charset:  config.ENV.DB_CHARSET,
	}

	DB = dbCfg.initMysqlDB()
}

// Opening the database connection
func (c *DbConfig) initMysqlDB() *gorm.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database, c.Charset)
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}

	db.AutoMigrate(&interfaces.Example{})

	return db
}
