package serviceQuote

import (
	"github.com/molodoymaxim/service-quotes/internal/repository/serviceQuote/psql"
	"github.com/molodoymaxim/service-quotes/internal/system"
)

// Встраиваем интерфейс PSQL
type ServiceQuote interface {
	psql.PSQL
}

// Встраиваем интерфейс PSQL в структуру
type serviceQuto struct {
	psql.PSQL
}

func New(sys *system.Systems) ServiceQuote {
	return serviceQuto{
		psql.New(sys.DB.PSQL),
	}
}
