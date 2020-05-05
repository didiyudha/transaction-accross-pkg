package main

import (
	"github.com/didiyudha/transaction-accross-pkg/config"
	"github.com/didiyudha/transaction-accross-pkg/internal/platform/postgres"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := config.Load(); err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("%+v\n", config.Cfg)
	_, err := postgres.Open(config.Cfg.DB)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("successfully connected to postgres")
}
