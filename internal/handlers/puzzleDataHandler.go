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

func (h *PuzzleDataHandler) GetPuzzleData(c *gin.Context) {
	id := c.Param("id")
	var puzzleData entity.PuzzleData

	if err := h.gorm.Where("id = ?", id).First(&puzzleData).Error; err != nil {
		h.logger.Errorf("Error fetching puzzle data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching puzzle data"})
		return
	}

	c.JSON(http.StatusOK, puzzleData)
}
