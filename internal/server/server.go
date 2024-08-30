package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nikurasuu/raetsel-backend/internal/config"
	"github.com/nikurasuu/raetsel-backend/internal/entity"
	"github.com/nikurasuu/raetsel-backend/internal/handlers"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	cfg    *config.Config
	logger *logrus.Logger
	db     *gorm.DB
}

func NewServer(cfg *config.Config, logger *logrus.Logger, db *gorm.DB) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}
}

func (s *Server) Start() error {
	r := gin.Default()

	s.db.AutoMigrate(&entity.PuzzleData{}, &entity.ResultData{})

	puzzleDataHandler := handlers.NewPuzzleDataHandler(s.logger, s.db)
	resultDataHandler := handlers.NewResultDataHandler(s.logger, s.db)

	addPuzzleDataRoutes(r, puzzleDataHandler)
	addResultDataRoutes(r, resultDataHandler)

	return r.Run(":" + fmt.Sprintf("%d", s.cfg.Server.Port))
}
