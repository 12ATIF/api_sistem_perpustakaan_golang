package usecase

import (
	"coba_dulu/internal/entity"
	"coba_dulu/internal/repository"
	"log"
)

type UserUsecase interface {
	CreateUser(user entity.User) (entity.User, error)
	Login(email string, password string) (interface{}, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}

func (uc *userUsecase) CreateUser(user entity.User) (entity.User, error) {
	return uc.userRepository.InsertUser(user)
}

func (uc *userUsecase) Login(email string, password string) (interface{}, error) {
	user, err := uc.userRepository.VerifyCredential(email, password)
	if err != nil {
		log.Println("Usecase Error:", err)
		return nil, err
	}
	return user, nil
}
