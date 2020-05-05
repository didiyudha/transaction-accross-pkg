package config

import (
	"github.com/didiyudha/transaction-accross-pkg/internal/platform/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Cfg Config

const (
	configFilename = "config"
	configType     = "yml"
	configPath     = "."
)

type Config struct {
	Port int `mapstructure:"port"`
	DB   postgres.Config `mapstructure:"db"`
}

func Load() error {
	viper.SetConfigName(configFilename)
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()
	viper.SetConfigType(configType)
	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Info("read configuration")
		return err
	}
	err := viper.Unmarshal(&Cfg)
	if err != nil {
		logrus.WithError(err).Info("unmarshal configuration")
		return err
	}
	return nil
}
