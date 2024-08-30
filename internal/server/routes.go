package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nikurasuu/raetsel-backend/internal/handlers"
)

func addPuzzleDataRoutes(r *gin.Engine, puzzleDataHandler *handlers.PuzzleDataHandler) {
	puzzleData := r.Group("/puzzle")

	puzzleData.POST("", puzzleDataHandler.PostPuzzleData)
}
