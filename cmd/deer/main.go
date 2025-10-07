package main

import (
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/reckedpr/deer/internal/api"
	"github.com/reckedpr/deer/internal/models"
	"github.com/reckedpr/deer/internal/util"
)

func main() {
	factData := models.Fact{
		FactPath: "./data/facts.json",
	}

	imageObject := models.Image{
		ImgPath: "./static/img",
	}

	var (
		files []string
		err   error
	)

	files, err = util.FilePathWalkDir(imageObject.ImgPath)
	if err != nil {
		log.Fatalf("failed to read images from %s\n%s", imageObject.ImgPath, err)
	}

	for _, file := range files {
		imageObject.ImgList = append(imageObject.ImgList, models.ImgJson{
			ImgURL: "/img/" + filepath.Base(file),
		})
	}

	api.ReadFacts(&factData)

	router := gin.Default()

	router.Static("/img", imageObject.ImgPath)

	router.GET("/", func(c *gin.Context) {
		api.ReturnImageJSON(c, &imageObject)
	})

	router.GET("/image", func(c *gin.Context) {
		api.ReturnImage(c, &imageObject)
	})

	router.GET("/fact", func(c *gin.Context) {
		api.ReturnFactJSON(c, &factData)
	})

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./web/public/favicon.ico")
	})

	router.Run("localhost:8080")
}
