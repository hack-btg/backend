package service

import (
	"context"

	"github.com/hack-btg/backend/banking-service/internal/clients/btg"
	"github.com/hack-btg/backend/banking-service/internal/domains/models"
	"github.com/hack-btg/backend/banking-service/storage"
)

type Bank interface {
	List(ctx context.Context) (models.Banks, error)
}

type BankService struct {
	bankRepo  BankRepository
	btgClient *btg.Client
}

func NewBankService() *BankService {
	return &BankService{
		bankRepo:  storage.NewBankRepo(),
		btgClient: btg.NewClient(),
	}
}

func (b *BankService) List(ctx context.Context) (bs models.Banks, err error) {
	empty := b.bankRepo.IsEmpty()
	if !empty {
		return b.bankRepo.List(), nil
	}
	bs, err = b.btgClient.GetBanks()
	if err != nil {
		return
	}
	b.bankRepo.Load(bs)
	return b.bankRepo.List(), nil
}
