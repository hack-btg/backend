package service

import (
	"context"

	"github.com/hack-btg/backend/banking-service/internal/domains/models"
	"github.com/hack-btg/backend/banking-service/storage"
)

type Social interface {
	Get(ctx context.Context) models.Social
}

type SocialService struct {
	socialRepo SocialRepository
}

func NewSocialService() *SocialService {
	return &SocialService{
		socialRepo: storage.NewSocialRepo(),
	}
}

func (ss *SocialService) Get(ctx context.Context) models.Social {
	return ss.socialRepo.Get()
}
