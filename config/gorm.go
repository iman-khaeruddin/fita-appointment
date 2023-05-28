//go:generate rm -fr mocks
//go:generate mockery --all

package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormMysql(dsn string) *gorm.DB {
	logLevel := 4 // 1 = Silent, 2 = Error, 3 = Warning, 4 = Info
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(logLevel)),
	})

	if err != nil {
		fmt.Println("database connection error")
	}
	return db
}
