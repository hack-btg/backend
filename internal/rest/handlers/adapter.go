package handlers

import (
	"github.com/hack-btg/backend/internal/jwt"
	"github.com/hack-btg/backend/internal/service"
)

type HTTPPrimaryAdapter struct {
	service       service.Orders
	tokenProvider jwt.TokenProvider
}

func NewHTTPPrimaryAdapter(s service.Orders, provider jwt.TokenProvider) *HTTPPrimaryAdapter {
	return &HTTPPrimaryAdapter{
		service:       s,
		tokenProvider: provider,
	}
}
