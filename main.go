package main

import (
	"fmt"
	"github.com/didiyudha/transaction-accross-pkg/config"
	"github.com/didiyudha/transaction-accross-pkg/domain/profile/handler"
	"github.com/didiyudha/transaction-accross-pkg/domain/profile/repository"
	"github.com/didiyudha/transaction-accross-pkg/domain/profile/usecase"
	userrepo "github.com/didiyudha/transaction-accross-pkg/domain/user/repository"
	"github.com/didiyudha/transaction-accross-pkg/internal/platform/postgres"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := config.Load(); err != nil {
		logrus.Fatal(err)
	}
	dbWrite, err := postgres.Open(config.Cfg.DB)
	if err != nil {
		logrus.Fatal(err)
	}
	dbRead, err := postgres.Open(config.Cfg.DB)
	if err != nil {
		logrus.Fatal(err)
	}
	userRepository := userrepo.NewUserRepository(dbWrite, dbRead)
	profileRepository := repository.NewProfileRepository(dbWrite, dbRead)
	profileUseCase := usecase.NewProfileUseCase(profileRepository, userRepository)
	profileHandler := handler.NewProfileHandler(profileUseCase)

	e := echo.New()
	e.POST("/profiles", profileHandler.CreateProfile)
	e.GET("/profiles/:id", profileHandler.FindByID)

	port := fmt.Sprintf(":%d", config.Cfg.Port)
	e.Logger.Fatal(e.Start(port))
}
