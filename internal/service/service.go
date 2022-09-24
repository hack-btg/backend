package service

import (
	"context"

	"github.com/hack-btg/backend/internal/domains/models"
	"github.com/hack-btg/backend/storage"
)

type Caixinha interface {
	Create(ctx context.Context, c models.Caixinha) (models.Caixinha, error)
	List(ctx context.Context) models.Caixinhas
	GetOne(ctx context.Context, id int) (models.Caixinha, error)
	Update(ctx context.Context, c models.Caixinha) error
	Delete(ctx context.Context, id int) error
}

type CXService struct {
	cxRepo CaixinhaRepository
}

func NewCXService() CXService {
	return CXService{
		cxRepo: storage.NewCaixinhaRepo(),
	}
}

func (s CXService) Create(ctx context.Context, c models.Caixinha) (cx models.Caixinha, err error) {
	return s.cxRepo.Create(c)
}

func (s CXService) GetOne(ctx context.Context, id int) (cx models.Caixinha, err error) {
	return s.cxRepo.GetOne(id)
}

func (s CXService) List(ctx context.Context) models.Caixinhas {
	return s.cxRepo.List()
}

func (s CXService) Update(ctx context.Context, c models.Caixinha) (err error) {
	return s.cxRepo.Update(c)
}

func (s CXService) Delete(ctx context.Context, id int) (err error) {
	return s.cxRepo.Delete(id)
}
