package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikurasuu/raetsel-backend/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ResultDataHandler struct {
	logger *logrus.Logger
	gorm   *gorm.DB
}

func NewResultDataHandler(logger *logrus.Logger, gorm *gorm.DB) *ResultDataHandler {
	return &ResultDataHandler{
		logger: logger,
		gorm:   gorm,
	}
}

func (h *ResultDataHandler) GetResultData(c *gin.Context) {
	id := c.Param("id")
	var resultData entity.ResultData

	if err := h.gorm.Where("id = ?", id).First(&resultData).Error; err != nil {
		h.logger.Errorf("Error fetching result data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching result data"})
		return
	}

	c.JSON(http.StatusOK, resultData)
}

func (h *ResultDataHandler) GetResultDataByPuzzleDataID(c *gin.Context) {
	id := c.Param("id")
	var resultData entity.ResultData

	if err := h.gorm.Where("puzzle_data_id = ?", id).First(&resultData).Error; err != nil {
		h.logger.Errorf("Error fetching result data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching result data"})
		return
	}

	c.JSON(http.StatusOK, resultData)
}
