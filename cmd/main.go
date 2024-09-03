package main

import (
	"strconv"

	"github.com/kamva/mgm/v3"
	"github.com/nikurasuu/raetsel-backend/internal/config"
	"github.com/nikurasuu/raetsel-backend/internal/server"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, _ := config.NewConfig()
	logger := logrus.New()

	err := mgm.SetDefaultConfig(nil, cfg.Mongo.DataBase, options.Client().ApplyURI("mongodb://"+cfg.Mongo.Host+":"+strconv.Itoa(cfg.Mongo.Port)))
	if err != nil {
		logger.Fatalf("Error setting up mgm: %v", err)
	}

	server := server.NewServer(cfg, logger)
	if err := server.Start(); err != nil {
		logger.Fatalf("Error starting the server: %v", err)
	}
}
