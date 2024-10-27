package service

import (
	"booking_service/internal/entity"
	"booking_service/internal/repository"
	"booking_service/pkg"
	"log"
)

type UserService struct {
	UserRepository repository.User
}

func NewUserService(repos *repository.Repository) *UserService {
	return &UserService{
		UserRepository: repos.User,
	}
}

func (s *UserService) CreateUser(user entity.User) (pkg.JWT, error) {
	var token pkg.JWT

	userID, err := s.UserRepository.CreateUser(user)
	if err != nil {
		log.Println(err)
		return token, err
	}

	token, err = pkg.CreateJWT(userID)
	if err != nil {
		log.Println(err)
		return token, err
	}

	return token, nil
}

func (s *UserService) Login(email, password string) (pkg.JWT, error) {
	var token pkg.JWT
	user, err := s.UserRepository.GetUserByLogIN(email, password)
	if err != nil {
		log.Println(err)
		return token, err
	}

	token, err = pkg.CreateJWT(user.ID)
	if err != nil {
		log.Println(err)
		return token, err
	}

	return token, nil
}

func (s *UserService) GetUserByToken(token string) (entity.User, error) {

	// Удаление Bearer из токена
	token = token[7:]

	userID, err := pkg.ValidateJWT(token)
	if err != nil {
		log.Println(err)
		return entity.User{}, err
	}

	user, err := s.UserRepository.GetUserByID(userID)
	if err != nil {
		log.Println(err)
		return entity.User{}, err
	}

	return user, nil
}
