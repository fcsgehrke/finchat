package services

import (
	"context"

	"github.com/fcsgehrke/finchat/internal/db/entities"
	"github.com/fcsgehrke/finchat/pkg/errors"
	"github.com/fcsgehrke/finchat/pkg/models"
)

func (s *Service) CreateUser(ctx context.Context, req *models.UserCreateRequest) (*models.UserCreateResponse, error) {
	pwd, err := s.crypt.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.CreateUser(ctx, &entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: pwd,
	})

	if err != nil {
		return nil, err
	}

	return &models.UserCreateResponse{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *Service) Login(ctx context.Context, req *models.UserLoginRequest) (*models.UserLoginResponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !s.crypt.ValidatePassword(req.Password, user.Password) {
		return nil, errors.ErrUserPasswordDoesntMatch
	}

	token, err := s.crypt.GenerateToken(user.Email, int(user.ID))
	if err != nil {
		return nil, err
	}

	return &models.UserLoginResponse{
		Token: token,
	}, nil
}
