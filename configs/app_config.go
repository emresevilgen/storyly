package configs

import (
	"github.com/spf13/viper"
)

type appConfig struct {
	v                       *viper.Viper
	GracefulShutdownTimeout int64             `required:"true" split_words:"true" yaml:"gracefulShutdownTimeout"`
	PostgreSql              PostgreSqlConfigs `required:"true" split_words:"true" yaml:"postgreSql"`
}

type PostgreSqlConfigs struct {
	Host string `required:"true" split_words:"true" yaml:"host"`
	Port string `required:"true" split_words:"true" yaml:"port"`
}

func (a *appConfig) readWithViper(shouldPanic bool) error {
	if a.v == nil {
		v := viper.New()
		v.AddConfigPath("./config")
		v.SetConfigName("application")
		v.SetConfigType("yaml")
		a.v = v
	}

	err := a.v.ReadInConfig()
	if err != nil {
		if shouldPanic {
			panic(err)
		}
		return err
	}

	err = a.v.Unmarshal(&AppConfig)
	if err != nil {
		if shouldPanic {
			panic(err)
		}
		return err
	}

	return nil
}
