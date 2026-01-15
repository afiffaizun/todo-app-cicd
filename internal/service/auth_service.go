package service

import (
	"errors"
	"github.com/afiffaizun/todo-app-cicd/internal/model"
	"github.com/afiffaizun/todo-app-cicd/internal/repository"
	"github.com/afiffaizun/todo-app-cicd/pkg/utils"
)

type AuthService interface {
	Register(name, email, password string) (*model.User, error)
	Login(email, password string) (string, *model.User, error)
	GetUserByID(id uint) (*model.User, error)
}

type authService struct {
	userRepo repository.UserRepository
	jwtUtil  *utils.JWTUtil
}

func NewAuthService(userRepo repository.UserRepository, jwtUtil *utils.JWTUtil) AuthService {
	return &authService{
		userRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

func (s *authService) Register(name, email, password string) (*model.User, error) {
	// Validate input
	if name == "" || email == "" || password == "" {
		return nil, errors.New("name, email, and password are required")
	}

	if len(password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	// Check if email already exists
	exists, err := s.userRepo.EmailExists(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email already registered")
	}

	// Create user
	user := &model.User{
		Name:  name,
		Email: email,
	}

	// Hash password
	if err := user.HashPassword(password); err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Save to database
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(email, password string) (string, *model.User, error) {
	// Validate input
	if email == "" || password == "" {
		return "", nil, errors.New("email and password are required")
	}

	// Find user by email
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	// Check password
	if !user.CheckPassword(password) {
		return "", nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := s.jwtUtil.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", nil, errors.New("failed to generate token")
	}

	return token, user, nil
}

func (s *authService) GetUserByID(id uint) (*model.User, error) {
	return s.userRepo.FindByID(id)
}