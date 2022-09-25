package service

import (
	"context"

	"github.com/hack-btg/backend/banking-service/internal/domains/models"
	"github.com/hack-btg/backend/banking-service/storage"
)

type User interface {
	Info(ctx context.Context) models.UserInfo
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: storage.NewUserRepo(),
	}
}

func (us *UserService) Info(ctx context.Context) models.UserInfo {
	return us.userRepo.Info()
}
