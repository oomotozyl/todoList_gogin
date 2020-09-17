package dao

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB DB
var (
	DB *gorm.DB
)

// InitPostgs InitPostgs
func InitPostgs() (err error) {
	dsn := "host=127.0.0.1 user=postgres password=password dbname=gogin port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) //???DB global,err local, how??? 解决：两个返回值时，可以提出返回err的方法
	if err != nil {
		return err
	}
	return
}
