package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nikurasuu/hms-solar-backend/internal/config"
	"github.com/nikurasuu/hms-solar-backend/internal/entity"
	"github.com/nikurasuu/hms-solar-backend/internal/handlers"
	"github.com/sirupsen/logrus"
)

type Server struct {
	cfg    *config.Config
	logger *logrus.Logger
}

func NewServer(cfg *config.Config, logger *logrus.Logger) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *Server) Start() error {
	r := gin.Default()

	solarData := entity.SolarData{}

	solarDataHandler := handlers.NewSolarDataHandler(&solarData)

	addSolarDataRoutes(r, solarDataHandler)

	return r.Run(":" + fmt.Sprintf("%d", s.cfg.Server.Port))
}
