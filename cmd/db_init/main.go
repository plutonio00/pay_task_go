package main

import (
	"github.com/plutonio00/pay-api/internal/config"
	"os/exec"
)

func main() {
	const configPath = "config/config"

	conf, err := config.Init(configPath)

	if err != nil {
		panic(err)
	}

	migrateInit(conf)
	// 	fixturesInit(conf)
}

func migrateInit(conf *config.Config) {
	cmd := exec.Command(
		"goose",
		"-dir",
		"internal/migration",
		"postgres",
		conf.Database.Postgres.DSN,
		"up",
	)

	err := cmd.Run()

	if err != nil {
		panic(err)
		return
	}
}
