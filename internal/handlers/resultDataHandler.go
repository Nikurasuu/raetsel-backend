package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"github.com/nikurasuu/raetsel-backend/internal/entity"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type ResultDataHandler struct {
	logger          *logrus.Logger
	mongoCollection *mgm.Collection
}

func NewResultDataHandler(logger *logrus.Logger, mongoCollection *mgm.Collection) *ResultDataHandler {
	return &ResultDataHandler{
		logger:          logger,
		mongoCollection: mongoCollection,
	}
}

func (h *ResultDataHandler) GetResultData(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf("Invalid id parameter: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	var resultData entity.ResultData
	filter := bson.M{"id": id}
	if err := h.mongoCollection.FindOne(mgm.Ctx(), filter).Decode(&resultData); err != nil {
		h.logger.Errorf("Error getting result data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting result data"})
		return
	}

	c.JSON(http.StatusOK, resultData)
}

func (h *ResultDataHandler) GetResultDataByPuzzleDataID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf("Invalid id parameter: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	var resultData entity.ResultData
	filter := bson.M{"puzzleDataId": id}
	if err := h.mongoCollection.FindOne(mgm.Ctx(), filter).Decode(&resultData); err != nil {
		h.logger.Errorf("Error getting result data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting result data"})
		return
	}

	c.JSON(http.StatusOK, resultData)
}
