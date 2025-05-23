package service

import (
	"github.com/molodoymaxim/service-quotes/internal/repository"
	"github.com/molodoymaxim/service-quotes/internal/service/serviceQuote"
)

type Service struct {
	ServiceQuote serviceQuote.QuoteService
}

func New(repo *repository.Repository) *Service {
	return &Service{
		ServiceQuote: serviceQuote.New(repo),
	}
}
