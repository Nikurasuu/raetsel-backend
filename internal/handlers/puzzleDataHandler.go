package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
	"github.com/nikurasuu/raetsel-backend/internal/entity"
	puzzlesolver "github.com/nikurasuu/raetsel-backend/internal/puzzleSolver"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type PuzzleDataHandler struct {
	logger            *logrus.Logger
	mongoCollection   *mgm.Collection
	puzzleSolver      *puzzlesolver.PuzzleSolver
	resultDataHandler *ResultDataHandler
}

func NewPuzzleDataHandler(logger *logrus.Logger, mongoCollection *mgm.Collection, puzzleSolver *puzzlesolver.PuzzleSolver, resultDataHandler *ResultDataHandler) *PuzzleDataHandler {
	return &PuzzleDataHandler{
		logger:            logger,
		mongoCollection:   mongoCollection,
		puzzleSolver:      puzzleSolver,
		resultDataHandler: resultDataHandler,
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

func (h *PuzzleDataHandler) PostPuzzleData(c *gin.Context) {
	var puzzleData entity.PuzzleData
	puzzleData.ID = uuid.New()
	puzzleData.CreatedAt = time.Now()
	puzzleData.UpdatedAt = time.Now()
	if err := c.ShouldBindJSON(&puzzleData); err != nil {
		h.logger.Errorf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.CreatePuzzleData(&puzzleData); err != nil {
		h.logger.Errorf("Error creating puzzle data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating puzzle data"})
		return
	}

	go func(puzzleData entity.PuzzleData) {
		resultData, err := h.puzzleSolver.SolvePuzzle(&puzzleData)
		if err != nil {
			h.logger.Errorf("Error solving puzzle: %v", err)
			return
		}
		if err := h.resultDataHandler.CreateResultData(&resultData); err != nil {
			h.logger.Errorf("Error creating result data: %v", err)
		}
	}(puzzleData)

	c.JSON(http.StatusCreated, puzzleData)
}

func (h *PuzzleDataHandler) CreatePuzzleData(puzzleData *entity.PuzzleData) error {
	if err := h.mongoCollection.Create(puzzleData); err != nil {
		h.logger.Errorf("Error creating puzzle data: %v", err)
		return err
	}
	h.logger.Info("Added puzzle data to database with id: ", puzzleData.ID)
	return nil
}
