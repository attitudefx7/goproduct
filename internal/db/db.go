package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func Init(c *Config) error {
	var err error

	fmt.Println(c.getDsn())

	db, err = gorm.Open(mysql.Open(c.getDsn()), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)

	return nil
}

func GetDB() *gorm.DB {
	return db
}