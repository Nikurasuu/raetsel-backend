package main

import (
	"github.com/nikurasuu/hms-solar-backend/internal/config"
	"github.com/nikurasuu/hms-solar-backend/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, _ := config.NewConfig()
	logger := logrus.New()

	server := server.NewServer(cfg, logger)
	if err := server.Start(); err != nil {
		logger.Fatalf("Error starting the server: %v", err)
	}
}
