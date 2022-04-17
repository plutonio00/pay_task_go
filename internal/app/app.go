package app

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"github.com/plutonio00/pay-api/internal/config"
	delivery "github.com/plutonio00/pay-api/internal/delivery/http"
	"github.com/plutonio00/pay-api/internal/repository"
	"github.com/plutonio00/pay-api/internal/server"
	"github.com/plutonio00/pay-api/internal/service"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) {
	conf, err := config.Init(configPath)

	if err != nil {
		logger.Error(err)
		return
	}

	db, err := sql.Open("postgres", conf.Database.Postgres.DSN)

	if err != nil {
		logger.Error(err)
		return
	}

	repos := repository.NewRepositories(db)

	services := service.NewServices(
		service.Deps{
			Repos: repos,
		},
	)

	handler := delivery.NewHandler(services)

	srv := server.NewServer(conf, handler.Init(conf))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {

		}
	}()

	logger.Info("server run")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
