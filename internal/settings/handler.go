package settings

import (
	"log"

	env "github.com/Netflix/go-env"
)

type Settings struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	Database struct {
		Name string
	}
}

var settings Settings

func init() {
	if _, err := env.UnmarshalFromEnviron(&settings); err != nil {
		log.Fatal("Error on initialize envs: ", err)
	}

}

func GetSettings() Settings {
	return settings
}
