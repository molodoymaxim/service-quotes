package handler

import (
	"github.com/molodoymaxim/service-quotes/internal/handler/http/serviceQuote"
	"github.com/molodoymaxim/service-quotes/internal/service"
	"time"
)

type Server struct {
	QuoteService *serviceQuote.QuoteServiceHandler
}

func New(serv *service.Service, timeLifeCtx int) *Server {
	t := time.Duration(timeLifeCtx) * time.Second
	return &Server{
		QuoteService: serviceQuote.New(serv.ServiceQuote, t),
	}
}
