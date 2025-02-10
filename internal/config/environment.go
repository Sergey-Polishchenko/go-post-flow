package config

import (
	"fmt"
	"sync"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/errors"
)

type Environment struct {
	Port string `env:"PORT" envDefault:"8080"`
	DB   *dbEnvironment
}

type dbEnvironment struct {
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Name     string `env:"DB_NAME,required"`
	Port     string `env:"DB_PORT,required"`
	Host     string `env:"DB_HOST,required"`
}

var (
	environment *Environment
	cfgErr      error
	once        sync.Once
)

func GetConfig(inmemory bool) (*Environment, error) {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			cfgErr = &errors.EnvLoadingError{Value: err}
		}

		environment = &Environment{}
		if !inmemory {
			environment.DB = &dbEnvironment{}
		}

		if err := env.Parse(environment); err != nil {
			cfgErr = errors.ErrEnvParsing
		}
	})
	return environment, cfgErr
}

func (e *dbEnvironment) ConnStr() string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		e.User,
		e.Password,
		e.Name,
		e.Host,
		e.Port,
	)
}
