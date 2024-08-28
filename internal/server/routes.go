package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nikurasuu/hms-solar-backend/internal/handlers"
)

func addSolarDataRoutes(r *gin.Engine, solarDataHandler *handlers.SolarDataHandler) {
	solarData := r.Group("/solardata")

	solarData.GET("", solarDataHandler.GetSolarData)
}
