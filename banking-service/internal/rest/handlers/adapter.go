package handlers

import (
	"github.com/hack-btg/backend/banking-service/internal/jwt"
	"github.com/hack-btg/backend/banking-service/internal/service"
)

type HTTPPrimaryAdapter struct {
	CaixinhaService service.Caixinha
	BankService     service.Bank
	tokenProvider   jwt.TokenProvider
}

func NewHTTPPrimaryAdapter(provider jwt.TokenProvider) *HTTPPrimaryAdapter {
	return &HTTPPrimaryAdapter{
		CaixinhaService: service.NewCXService(),
		BankService:     service.NewBankService(),
		tokenProvider:   provider,
	}
}
