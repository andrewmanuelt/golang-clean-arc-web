package repository

import "golang-clean-arc-web/entity"

type ExampleRepository interface {
	Get(app_ver string) *entity.App
}
