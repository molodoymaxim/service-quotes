package serviceQuote

import (
	"github.com/molodoymaxim/service-quotes/internal/service/serviceQuote"
	"time"
)

type QuoteServiceHandler struct {
	quotaService serviceQuote.QuoteService
	timeLifeCtx  time.Duration
}

func New(quotaService serviceQuote.QuoteService, timeLifeCtx time.Duration) *QuoteServiceHandler {
	return &QuoteServiceHandler{
		quotaService: quotaService,
		timeLifeCtx:  timeLifeCtx,
	}
}
