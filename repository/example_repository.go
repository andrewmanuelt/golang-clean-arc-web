package repository

import (
	"golang-clean-arc-web/entity"

	"gorm.io/gorm"
)

type ExampleRepository interface {
	Select() (response []entity.App)
	Insert(request entity.App)
	Update(app_ver string, update_data interface{})
	Delete(app_ver string)

	GetLast() (response entity.App)
}

type exampleRepositoryImpl struct {
	Database *gorm.DB
}

func (repository *exampleRepositoryImpl) Select() (response []entity.App) {
	var app []entity.App

	repository.Database.Find(&app)

	for _, row := range app {
		response = append(response, entity.App{
			AppName: row.AppName,
			AppVer:  row.AppVer,
		})
	}

	return response
}

func (repository *exampleRepositoryImpl) Insert(request entity.App) {
	app := entity.App{
		AppName: request.AppName,
		AppVer:  request.AppVer,
	}

	repository.Database.Create(&app)
}

func (repository *exampleRepositoryImpl) Update(app_ver string, update_data interface{}) {
	var app []entity.App

	repository.Database.Model(&app).Where("app_ver", app_ver).Updates(update_data)
}

func (repository *exampleRepositoryImpl) Delete(app_ver string) {
	var app []entity.App

	repository.Database.Where("app_ver", app_ver).Delete(&app)

	repository.Database.Find(&app)
}

func (repository *exampleRepositoryImpl) GetLast() (response entity.App) {
	var app entity.App

	repository.Database.Last(&app)

	return entity.App{
		AppName: app.AppName,
		AppVer:  app.AppVer,
	}
}

func NewExampleRepository(database *gorm.DB) ExampleRepository {
	return &exampleRepositoryImpl{
		Database: database,
	}
}
