package storage

import "github.com/hack-btg/backend/banking-service/internal/domains/models"

type BankStorage struct {
	store models.Banks
}

func NewBankRepo() *BankStorage {
	return &BankStorage{
		store: models.Banks{},
	}
}

func (b *BankStorage) Load(bs models.Banks) {
	b.store = bs
}

func (b *BankStorage) List() models.Banks {
	return b.store
}

func (b *BankStorage) IsEmpty() bool {
	return len(b.store) == 0
}
