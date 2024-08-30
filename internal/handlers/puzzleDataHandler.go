package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikurasuu/raetsel-backend/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PuzzleDataHandler struct {
	logger *logrus.Logger
	gorm   *gorm.DB
}

func NewPuzzleDataHandler(logger *logrus.Logger, gorm *gorm.DB) *PuzzleDataHandler {
	return &PuzzleDataHandler{
		logger: logger,
		gorm:   gorm,
	}
}

func (h *PuzzleDataHandler) PostPuzzleData(c *gin.Context) {
	var puzzleData entity.PuzzleData
	if err := c.BindJSON(&puzzleData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Infof("Received puzzle data: %+v", puzzleData)

	// TODO: Validate the puzzle data

	// TODO: Save the puzzle data to the database

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
