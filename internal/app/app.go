package app

import (
	"github.com/molodoymaxim/service-quotes/internal/config"
	handler "github.com/molodoymaxim/service-quotes/internal/handler"
	"github.com/molodoymaxim/service-quotes/internal/repository"
	"github.com/molodoymaxim/service-quotes/internal/server"
	routerHTTP "github.com/molodoymaxim/service-quotes/internal/server/http"
	"github.com/molodoymaxim/service-quotes/internal/service"
	"github.com/molodoymaxim/service-quotes/internal/system"
	"github.com/molodoymaxim/service-quotes/internal/types"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Start() {

	cfgHTTP := &types.ConfigHTTP{}
	cfgPostgres := &types.ConfigPostgres{}
	cfgConnPostgres := &types.ConfigConnPostgres{}

	err := config.GetConfigsENV("./", ".env", []any{
		cfgHTTP,
		cfgPostgres,
		cfgConnPostgres,
	})
	if err != nil {
		log.Fatal(err)
	}

	syst, err := system.New(
		cfgConnPostgres,
		cfgPostgres,
	)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New(syst)

	serv := service.New(repo)

	srvHTTP := server.New(
		cfgHTTP.Port,
	)

	routerHTTP := routerHTTP.New(
		handler.New(
			serv,
			cfgHTTP.TimeLifeCtx,
		),
	)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	if err := srvHTTP.Start(c, routerHTTP.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
