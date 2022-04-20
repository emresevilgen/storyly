package configs

import (
	"github.com/fsnotify/fsnotify"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
)

var (
	AppConfig appConfig
	Secrets   secrets
)

func InitConfigs() {
	if os.Getenv("CONFIG_ENV_ENABLED") == "true" {
		envconfig.MustProcess("", &AppConfig)
		envconfig.MustProcess("", &Secrets)
		return
	}

	AppConfig.readWithViper(true)
	AppConfig.v.WatchConfig()
	AppConfig.v.OnConfigChange(func(in fsnotify.Event) {
		err := AppConfig.readWithViper(false)
		if err != nil {
			log.Println("Error on refreshing application configs due to file change, error: ", err)
			return
		}
		log.Println("Application configs are changed")
	})

	Secrets.readWithViper(true)
	Secrets.v.WatchConfig()
	Secrets.v.OnConfigChange(func(in fsnotify.Event) {
		err := Secrets.readWithViper(false)
		if err != nil {
			log.Println("Error on refreshing secrets due to file change, error: ", err)
			return
		}
		log.Println("Secret configs are changed")
	})
}
