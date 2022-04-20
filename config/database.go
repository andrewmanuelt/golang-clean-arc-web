package config

import (
	"golang-clean-arc/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dsn := datasource()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	helper.ShowError(err)

	return db
}

func datasource() string {
	dsn := []string{
		Env("DB_USERNAME"),
		":",
		Env("DB_PASSWORD"),
		"@tcp(",
		Env("DB_URL"),
		":",
		Env("DB_PORT"),
		")/",
		Env("DB_NAME"),
		"?parseTime=true",
	}

	return helper.AppendString(dsn)
}
