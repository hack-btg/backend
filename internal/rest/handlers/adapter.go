package handlers

import (
	"github.com/hack-btg/backend/internal/jwt"
	"github.com/hack-btg/backend/internal/service"
)

type HTTPPrimaryAdapter struct {
	CaixinhaService service.Caixinha
	tokenProvider   jwt.TokenProvider
}

func NewHTTPPrimaryAdapter(provider jwt.TokenProvider) *HTTPPrimaryAdapter {
	return &HTTPPrimaryAdapter{
		CaixinhaService: service.NewCXService(),
		tokenProvider:   provider,
	}
}
