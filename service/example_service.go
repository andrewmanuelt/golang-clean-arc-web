package service

import (
	"golang-clean-arc-web/entitiy"
	"golang-clean-arc-web/model"
	"golang-clean-arc-web/repository"
)

type ExampleService interface {
	Get() (response []model.ExampleResponse)
	Create(appName string, appVer string) (response model.ExampleResponse)
	Modify(app_ver string, update_data interface{}) (response []model.ExampleResponse)
	Erase(app_ver string) (response []model.ExampleResponse)
}

type exampleServiceImpl struct {
	ExampleRepository repository.ExampleRepository
}

func (example *exampleServiceImpl) Get() (response []model.ExampleResponse) {
	data := example.ExampleRepository.Select()

	for _, row := range data {
		response = append(response, model.ExampleResponse{
			AppName: row.AppName,
			AppVer:  row.AppVer,
		})
	}

	return response
}

func (example *exampleServiceImpl) Create(appName string, appVer string) (response model.ExampleResponse) {
	example.ExampleRepository.Insert(entitiy.App{
		AppName: appName,
		AppVer:  appVer,
	})

	app := example.ExampleRepository.GetLast()

	return model.ExampleResponse{
		AppName: app.AppName,
		AppVer:  app.AppVer,
	}
}

func (example *exampleServiceImpl) Modify(app_ver string, update_data interface{}) (response []model.ExampleResponse) {
	example.ExampleRepository.Update(app_ver, update_data)

	data := example.ExampleRepository.Select()

	for _, row := range data {
		response = append(response, model.ExampleResponse{
			AppName: row.AppName,
			AppVer:  row.AppVer,
		})
	}

	return response
}

func (example *exampleServiceImpl) Erase(app_ver string) (response []model.ExampleResponse) {

	example.ExampleRepository.Delete(app_ver)

	data := example.ExampleRepository.Select()

	for _, row := range data {
		response = append(response, model.ExampleResponse{
			AppName: row.AppName,
			AppVer:  row.AppVer,
		})
	}

	return response
}

func NewExampleService(repository *repository.ExampleRepository) ExampleService {
	return &exampleServiceImpl{
		ExampleRepository: *repository,
	}
}
