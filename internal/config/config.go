package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type (
	Config struct {
		Server   ServerConfig
		Database DatabaseConfig
	}

	ServerConfig struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	}

	DatabaseConfig struct {
		Postgres PostgresConfig
	}

	PostgresConfig struct {
		User     string
		Password string
		Host     string
		Port     string
		DBName   string
		DSN      string
	}
)

func Init(configPath string) (*Config, error) {

	if err := parseConfigFile(configPath); err != nil {
		return nil, err
	}

	if err := parseEnvFile(); err != nil {
		return nil, err
	}

	var conf Config

	if err := unmarshal(&conf); err != nil {
		return nil, err
	}

	setFromEnvFile(&conf)

	return &conf, nil
}

func parseConfigFile(filepath string) error {
	path := strings.Split(filepath, "/")

	viper.AddConfigPath(path[0])
	viper.SetConfigName(path[1])

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetConfigName("env")
	viper.MergeInConfig()

	return nil
}

func unmarshal(cfg *Config) error {

	if err := viper.UnmarshalKey("server", &cfg.Server); err != nil {
		return err
	}

	return nil
}

func setFromEnvFile(conf *Config) {
	conf.Database.Postgres.User = os.Getenv("POSTGRES_USER")
	conf.Database.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	conf.Database.Postgres.Host = os.Getenv("POSTGRES_HOST")
	conf.Database.Postgres.Port = os.Getenv("POSTGRES_PORT")
	conf.Database.Postgres.DBName = os.Getenv("POSTGRES_DB")
	conf.Database.Postgres.DSN = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Database.Postgres.Host,
		conf.Database.Postgres.Port,
		conf.Database.Postgres.User,
		conf.Database.Postgres.Password,
		conf.Database.Postgres.DBName)
}

func parseEnvFile() error {

	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}
