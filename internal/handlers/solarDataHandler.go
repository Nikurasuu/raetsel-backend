package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikurasuu/hms-solar-backend/internal/entity"
)

type SolarDataHandler struct {
	solarData *entity.SolarData
}

func NewSolarDataHandler(solarData *entity.SolarData) *SolarDataHandler {
	return &SolarDataHandler{
		solarData: solarData,
	}
}

func (h *SolarDataHandler) GetSolarData(c *gin.Context) {
	c.JSON(http.StatusOK, h.solarData)
}
