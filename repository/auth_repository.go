package repository

import (
	"golang-clean-arc-web/entity"
	"time"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(email string) (user entity.User)
	Register(user entity.User)
}

type authRepositoryImpl struct {
	Database *gorm.DB
}

func (repository *authRepositoryImpl) Register(user entity.User) {
	repository.Database.Create(&user)
}

func (repository *authRepositoryImpl) Login(email string) (user entity.User) {

	tx := repository.Database.Session(&gorm.Session{PrepareStmt: true})

	tx.Where("email", email).Last(&user)

	return entity.User{
		Username:   user.Username,
		Password:   user.Password,
		Email:      user.Email,
		Last_login: time.Now(),
	}
}

func NewAuthRepository(database *gorm.DB) AuthRepository {
	return &authRepositoryImpl{
		Database: database,
	}
}
