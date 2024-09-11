package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
	"github.com/nikurasuu/raetsel-backend/internal/entity"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type PuzzleDataHandler struct {
	logger          *logrus.Logger
	mongoCollection *mgm.Collection
}

func NewPuzzleDataHandler(logger *logrus.Logger, mongoCollection *mgm.Collection) *PuzzleDataHandler {
	return &PuzzleDataHandler{
		logger:          logger,
		mongoCollection: mongoCollection,
	}
}

func (h *PuzzleDataHandler) GetPuzzleData(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		h.logger.Errorf("Invalid id parameter: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	var puzzleData entity.PuzzleData
	filter := bson.M{"id": id}
	if err := h.mongoCollection.FindOne(mgm.Ctx(), filter).Decode(&puzzleData); err != nil {
		h.logger.Errorf("Error getting puzzle data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting puzzle data"})
		return
	}
	c.JSON(http.StatusOK, puzzleData)
}
