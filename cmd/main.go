package main

import (
	"github.com/nikurasuu/raetsel-backend/internal/config"
	"github.com/nikurasuu/raetsel-backend/internal/server"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg, _ := config.NewConfig()
	logger := logrus.New()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Error opening the database: %v", err)
	}

	server := server.NewServer(cfg, logger, db)
	if err := server.Start(); err != nil {
		logger.Fatalf("Error starting the server: %v", err)
	}
}
