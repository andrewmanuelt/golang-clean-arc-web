package web

import (
	"golang-clean-arc-web/entity"
	webModel "golang-clean-arc-web/model/web"
	"golang-clean-arc-web/repository"
)

type AuthService interface {
	Login(email string) (response webModel.LoginResponse)
	Register(request webModel.RegisterRequest)
}

type authServiceImpl struct {
	AuthRepository repository.AuthRepository
}

func (authService *authServiceImpl) Register(request webModel.RegisterRequest) {
	authService.AuthRepository.Register(entity.User{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	})
}

func (authService *authServiceImpl) Login(email string) (response webModel.LoginResponse) {
	data := authService.AuthRepository.Login(email)

	return webModel.LoginResponse{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}
}

func NewAuthService(authRepository *repository.AuthRepository) AuthService {
	return &authServiceImpl{
		AuthRepository: *authRepository,
	}
}
