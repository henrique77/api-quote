package database

import (
	"log"
	"time"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/henrique77/api-quote/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(env *config.Env) *gorm.DB {
	env = config.ReadEnvs()

	config := mysqlDriver.Config{
		User:                 env.DBUserName,
		Passwd:               env.DBPassword,
		DBName:               env.DBName,
		Net:                  "tcp",
		Addr:                 env.DBHost + ":" + env.DBPort,
		Loc:                  time.UTC,
		AllowNativePasswords: true,
	}

	db, err := gorm.Open(mysql.Open(config.FormatDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	return db
}
