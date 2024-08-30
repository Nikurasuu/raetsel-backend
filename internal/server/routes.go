package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nikurasuu/raetsel-backend/internal/handlers"
)

func addPuzzleDataRoutes(r *gin.Engine, puzzleDataHandler *handlers.PuzzleDataHandler) {
	puzzleData := r.Group("/puzzle")

	puzzleData.GET("/:id", puzzleDataHandler.GetPuzzleData)
}

func addResultDataRoutes(r *gin.Engine, resultDataHandler *handlers.ResultDataHandler) {
	resultData := r.Group("/result")

	resultData.GET("/:id", resultDataHandler.GetResultData)
	resultData.GET("/puzzle/:id", resultDataHandler.GetResultDataByPuzzleDataID)
}
