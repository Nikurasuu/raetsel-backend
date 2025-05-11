package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"github.com/nikurasuu/raetsel-backend/internal/config"
	"github.com/nikurasuu/raetsel-backend/internal/entity"
	"github.com/nikurasuu/raetsel-backend/internal/handlers"
	puzzlesolver "github.com/nikurasuu/raetsel-backend/internal/puzzleSolver"
	"github.com/nikurasuu/raetsel-backend/internal/wordlist"
	"github.com/sirupsen/logrus"
)

type Server struct {
	cfg      *config.Config
	logger   *logrus.Logger
	wordlist *wordlist.WordList
}

func NewServer(cfg *config.Config, logger *logrus.Logger, wordlist *wordlist.WordList) *Server {
	return &Server{
		cfg:      cfg,
		logger:   logger,
		wordlist: wordlist,
	}
}

func (s *Server) Start() error {
	r := gin.Default()

	puzzleDataCollection := mgm.Coll(&entity.PuzzleData{})
	resultDataCollection := mgm.Coll(&entity.ResultData{})

	puzzlesolver := puzzlesolver.NewPuzzleSolver(s.logger, s.wordlist)

	resultDataHandler := handlers.NewResultDataHandler(s.logger, resultDataCollection)
	puzzleDataHandler := handlers.NewPuzzleDataHandler(s.logger, puzzleDataCollection, puzzlesolver, resultDataHandler)

	addPuzzleDataRoutes(r, puzzleDataHandler)
	addResultDataRoutes(r, resultDataHandler)

	return r.Run(":" + fmt.Sprintf("%d", s.cfg.Server.Port))
}
