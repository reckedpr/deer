package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/reckedpr/deer/internal/middleware"
	"github.com/reckedpr/deer/internal/models"
	"go.uber.org/zap"
)

func ReturnFactJSON(c *gin.Context, factObject *models.Fact) {
	file, err := os.Stat(factObject.FactPath)
	if err != nil {
		zap.S().Errorw("Unable to get FileInfo", "filepath", factObject.FactPath)
	}

	if file.Size() != factObject.FactFileSize {
		zap.S().Infow("Updating fact list", "file", factObject.FactPath)
		err = ReadFacts(factObject)
		if err != nil {
			ReturnError(c, http.StatusInternalServerError, "failed to refetch random facts", err)
			zap.S().Errorw("Failed to read facts", "file", factObject.FactPath, "error", err)
			return
		}
		factObject.FactFileSize = file.Size()
	}

	randomIndex := rand.Intn(len(factObject.FactList))
	chosenFact := factObject.FactList[randomIndex]

	middleware.NoCacheHeaders(c)
	c.IndentedJSON(http.StatusOK, chosenFact)
}

func ReadFacts(factObject *models.Fact) error {
	factFile, err := os.ReadFile(factObject.FactPath)
	if err != nil {
		zap.S().Errorw("Failed to ReadFile", "file", factObject.FactPath, "error", err)
		return err
	}

	err = json.Unmarshal(factFile, &factObject.FactList)
	if err != nil {
		zap.S().Errorw("Failed to unmarhsal json", "file", factObject.FactPath, "error", err)
		return err
	}

	zap.S().Infow("Fact file loaded", "file", factObject.FactPath, "amount", len(factObject.FactList))
	return nil
}
