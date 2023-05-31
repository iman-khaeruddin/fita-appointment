package test_util

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// NewQueryDBMock creates mockQuery and mockDB
func NewQueryDBMock() (sqlmock.Sqlmock, *gorm.DB, bool) {
	mockConn, mockQuery, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Println(err)
		return nil, nil, false
	}
	mockDB, err := GormMysqlFromConnMock(mockConn)
	if err != nil {
		log.Println(err)
		return nil, nil, false
	}
	return mockQuery, mockDB, true
}

// GormMysqlFromConnMock creates gorm connection for testing purpose
func GormMysqlFromConnMock(db *sql.DB) (*gorm.DB, error) {
	mysqlConfig := mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}
	options := &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Time{}
		},
	}

	return gorm.Open(mysql.New(mysqlConfig), options)
}
