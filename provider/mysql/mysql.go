package ProviderMysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/config"
)

func NewMysql(cfg *config.Config) (*gorm.DB, error) {
	//dsn := os.Getenv("DOMAIN")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.Mysql.User, cfg.Mysql.Password,
		cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
