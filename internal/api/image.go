package api

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reckedpr/deer/internal/middleware"
	"github.com/reckedpr/deer/internal/models"
)

func ReturnImageJSON(c *gin.Context, imgObject *models.Image) {
	chosenImg := randomImage(imgObject)

	scheme := c.Request.Header.Get("X-Forwarded-Proto")
	if scheme == "" {
		if c.Request.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}
	}

	host := c.Request.Host
	chosenImg.ImgURL = scheme + "://" + host + chosenImg.ImgURL

	middleware.NoCacheHeaders(c)
	c.IndentedJSON(http.StatusOK, chosenImg)
}

func ReturnImage(c *gin.Context, imgObject *models.Image) {
	chosenImg := randomImage(imgObject)

	middleware.NoCacheHeaders(c)
	c.File("./static/" + chosenImg.ImgURL)
}

func randomImage(imgObject *models.Image) models.ImgJson {
	randomIndex := rand.Intn(len(imgObject.ImgList))
	return imgObject.ImgList[randomIndex]
}
