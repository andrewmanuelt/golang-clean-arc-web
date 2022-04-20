package repository

import (
	"golang-clean-arc/entitiy"

	"gorm.io/gorm"
)

type ExampleRepository interface {
	Select() (response []entitiy.App)
	Insert(request entitiy.App)
	Update(app_ver string, update_data interface{})
	Delete(app_ver string)

	GetLast() (response entitiy.App)
}

type exampleRepositoryImpl struct {
	Database *gorm.DB
}

func (repository *exampleRepositoryImpl) Select() (response []entitiy.App) {
	var app []entitiy.App

	repository.Database.Find(&app)

	for _, row := range app {
		response = append(response, entitiy.App{
			AppName: row.AppName,
			AppVer:  row.AppVer,
		})
	}

	return response
}

func (repository *exampleRepositoryImpl) Insert(request entitiy.App) {
	app := entitiy.App{
		AppName: request.AppName,
		AppVer:  request.AppVer,
	}

	repository.Database.Create(&app)
}

func (repository *exampleRepositoryImpl) Update(app_ver string, update_data interface{}) {
	var app []entitiy.App

	repository.Database.Model(&app).Where("app_ver", app_ver).Updates(update_data)
}

func (repository *exampleRepositoryImpl) Delete(app_ver string) {
	var app []entitiy.App

	repository.Database.Where("app_ver", app_ver).Delete(&app)

	repository.Database.Find(&app)
}

func (repository *exampleRepositoryImpl) GetLast() (response entitiy.App) {
	var app entitiy.App

	repository.Database.Last(&app)

	return entitiy.App{
		AppName: app.AppName,
		AppVer:  app.AppVer,
	}
}

func NewExampleRepository(database *gorm.DB) ExampleRepository {
	return &exampleRepositoryImpl{
		Database: database,
	}
}
