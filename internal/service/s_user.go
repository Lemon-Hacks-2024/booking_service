package service

import (
	"booking_service/internal/entity"
	"booking_service/internal/repository"
	"booking_service/pkg"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

	// Хеширование пароля
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		log.Println(err)
		return token, err
	}
	user.Password = hashedPassword

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

func hashPassword(password string) (string, error) {
	// Используем bcrypt для генерации хеша пароля
	// Второй параметр - cost factor, который определяет сложность хеширования
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hash), nil
}

func (s *UserService) Login(email, password string) (pkg.JWT, error) {
	var token pkg.JWT

	// Хеширование пароля
	hashedPassword, err := hashPassword(password)
	if err != nil {
		log.Println(err)
		return token, err
	}
	password = hashedPassword

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
