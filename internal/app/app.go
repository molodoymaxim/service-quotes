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

	// Конфигурации
	cfgHTTP := &types.ConfigHTTP{}
	cfgPostgres := &types.ConfigPostgres{}
	cfgConnPostgres := &types.ConfigConnPostgres{}

	// Подгружаем конфигурации
	err := config.GetConfigsENV("./", ".env", []any{
		cfgHTTP,
		cfgPostgres,
		cfgConnPostgres,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Подключаем все зависимости
	syst, err := system.New(
		cfgConnPostgres,
		cfgPostgres,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем репозиторий
	repo := repository.New(syst)

	// Создвние сервисов
	serv := service.New(repo)

	// Создаем HTTP сервер
	srvHTTP := server.New(
		cfgHTTP.Port,
	)

	// Создаем HTTP роутер
	routerHTTP := routerHTTP.New(
		handler.New(
			serv,
			cfgHTTP.TimeLifeCtx,
		),
	)

	// Канал завершения
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// Запускаем HTTP сервер
	if err := srvHTTP.Start(c, routerHTTP.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
