package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/torrez/pkg"
	"github.com/torrez/src/dtos"
	"github.com/torrez/src/models"
	"github.com/torrez/src/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) CreateUser(req dtos.CreateUserRequest) (*dtos.UserResponse, error) {
	id := uuid.New()
	hashedPassword, err := pkg.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Phone:    req.Phone,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &dtos.UserResponse{
		ID:        id.String(),
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Rol:       string(user.Rol),
		Slug:      user.Slug,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil

}

func (s *UserService) Login(req dtos.LoginRequest) (*dtos.LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if !pkg.CheckPasswordHash(req.Password, user.Password) {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Incluir el rol en el token JWT
	token, err := pkg.GenerateJWT(user.ID.String(), string(user.Rol))
	if err != nil {
		return nil, err
	}

	return &dtos.LoginResponse{
		Token: token,
	}, nil
}

func (s *UserService) CreateAdmin(req dtos.CreateUserRequest) (*dtos.UserResponse, error) {
	id := uuid.New()
	hashedPassword, err := pkg.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Phone:    req.Phone,
		Rol:      "Administrador",
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &dtos.UserResponse{
		ID:        id.String(),
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Rol:       string(user.Rol),
		Slug:      user.Slug,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil
}
