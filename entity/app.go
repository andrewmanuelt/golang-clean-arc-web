package entity

import "gorm.io/gorm"

type App struct {
	gorm.Model
	AppName string
	AppVer  string
}
