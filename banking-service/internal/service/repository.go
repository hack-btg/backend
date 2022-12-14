package service

import "github.com/hack-btg/backend/banking-service/internal/domains/models"

type CaixinhaRepository interface {
	Create(cx models.Caixinha) (models.Caixinha, error)
	List() models.Caixinhas
	GetOne(id int) (models.Caixinha, error)
	Delete(id int) error
	Update(cx models.Caixinha) error
}

type BankRepository interface {
	List() models.Banks
	IsEmpty() bool
	Load(models.Banks)
}

type UserRepository interface {
	Info() models.UserInfo
}

type SocialRepository interface {
	Get() models.Social
}
