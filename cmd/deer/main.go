package main

import (
	"fmt"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/reckedpr/deer/internal/api"
	"github.com/reckedpr/deer/internal/middleware"
	"github.com/reckedpr/deer/internal/models"
	"github.com/reckedpr/deer/internal/util"
	"go.uber.org/zap"
)

func init() {
	logger := zap.Must(zap.NewProduction())
	zap.ReplaceGlobals(logger)
}

func main() {
	const port int = 8080

	listenAddr := fmt.Sprintf(":%d", port)

	factObject := models.Fact{
		FactPath: "./data/facts.json",
	}

	imageObject := models.Image{
		ImgPath: "./static/img",
	}

	zap.S().Infow("Starting Deer API Server")

	files, err := util.FilePathWalkDir(imageObject.ImgPath)
	if err != nil {
		zap.S().Fatalw("Failed to read images", "path", imageObject.ImgPath, "err", err)
	}

	for _, file := range files {
		imageObject.ImgList = append(imageObject.ImgList, models.ImgJson{
			ImgURL: "/img/" + filepath.Base(file),
		})
	}

	zap.S().Infow("Loaded images", "path", imageObject.ImgPath, "amount", len(imageObject.ImgList))

	api.ReadFacts(&factObject)

	// init
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.SetTrustedProxies(nil)

	router.Use(
		middleware.GinZapMiddleware(),
		gin.Recovery(),
		cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET"},
			AllowHeaders: []string{"Origin"},
		}),
	)

	// qol
	router.Static("/img", imageObject.ImgPath)

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./web/public/favicon.ico")
	})

	// endpoints
	router.GET("/", func(c *gin.Context) {
		api.ReturnImageJSON(c, &imageObject)
	})

	router.GET("/image", func(c *gin.Context) {
		api.ReturnImage(c, &imageObject)
	})

	router.GET("/fact", func(c *gin.Context) {
		api.ReturnFactJSON(c, &factObject)
	})

	zap.S().Infow("Router starting", "addr", listenAddr)
	router.Run(listenAddr)
}
