package configs

import "github.com/spf13/viper"

type secrets struct {
	v                  *viper.Viper
	PostgreSqlUser     string `required:"true" split_words:"true" json:"postgreSqlUser"`
	PostgreSqlPassword string `required:"true" split_words:"true" json:"postgreSqlPassword"`
}

func (s *secrets) readWithViper(shouldPanic bool) error {
	if s.v == nil {
		v := viper.New()
		v.AddConfigPath("./config")
		v.SetConfigName("secrets")
		v.SetConfigType("json")
		s.v = v
	}

	err := s.v.ReadInConfig()
	if err != nil {
		if shouldPanic {
			panic(err)
		}
		return err
	}

	err = s.v.Unmarshal(&Secrets)
	if err != nil {
		if shouldPanic {
			panic(err)
		}
		return err
	}

	return nil
}
