package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/reckedpr/deer/internal/models"
)

func ReturnFactJSON(c *gin.Context, factData *models.Fact) {
	err := ReadFacts(factData)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, "failed to refetch random facts", err)
		return
	}

	randomIndex := rand.Intn(len(factData.FactList))
	chosenFact := factData.FactList[randomIndex]

	c.IndentedJSON(http.StatusOK, chosenFact)
}

func ReadFacts(factData *models.Fact) error {
	factFile, err := os.ReadFile(factData.FactPath)
	if err != nil {
		return fmt.Errorf("failed to read fact path @ %s\n\n%s", factData.FactPath, err)
	}

	err = json.Unmarshal(factFile, &factData.FactList)
	if err != nil {
		return fmt.Errorf("failed to open %s\n%s", factFile, err)
	}

	return nil
}
