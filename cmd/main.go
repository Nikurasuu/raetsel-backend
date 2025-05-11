package main

import (
	"strconv"

	"github.com/kamva/mgm/v3"
	"github.com/nikurasuu/raetsel-backend/internal/config"
	"github.com/nikurasuu/raetsel-backend/internal/server"
	"github.com/nikurasuu/raetsel-backend/internal/wordlist"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, _ := config.NewConfig()
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	wordlist, wordlistErr := wordlist.NewWordList("words.db")
	if wordlistErr != nil {
		logger.Fatalf("Error creating wordlist: %v", wordlistErr)
	}

	mgmErr := mgm.SetDefaultConfig(nil, cfg.Mongo.DataBase, options.Client().ApplyURI("mongodb://"+cfg.Mongo.Host+":"+strconv.Itoa(cfg.Mongo.Port)))
	if mgmErr != nil {
		logger.Fatalf("Error setting up mgm: %v", mgmErr)
	}

	server := server.NewServer(cfg, logger, wordlist)
	if err := server.Start(); err != nil {
		logger.Fatalf("Error starting the server: %v", err)
	}
}
