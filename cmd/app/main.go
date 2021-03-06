package main

import (
	_ "github.com/plutonio00/pay-api/docs"
	"github.com/plutonio00/pay-api/internal/app"
)

const configPath = "config/config"

// @title Pay Task API
// @version 1.0

// @host 127.0.0.1:8080
// @BasePath /api/v1/
func main() {
	app.Run(configPath)
}
