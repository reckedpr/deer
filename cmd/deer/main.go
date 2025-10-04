package main

import (
	"math/rand"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var deerList []Deer

type Deer struct {
	ID     int    `json:"id"`
	ImgURL string `json:"img_url,omitempty"`
}

// proudly stolen from stackoverflow !
func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func getDeer(c *gin.Context) {
	randomIndex := rand.Intn(len(deerList))
	deer := deerList[randomIndex]

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	host := c.Request.Host
	deer.ImgURL = scheme + "://" + host + deer.ImgURL

	c.IndentedJSON(http.StatusOK, deer)

}

func main() {
	var (
		err   error
		files []string
	)

	imgDir := "./static/img"

	files, err = FilePathWalkDir(imgDir)
	if err != nil {
		panic(err)
	}

	for i, file := range files {
		deerList = append(deerList, Deer{
			ID:     i + 1,
			ImgURL: "/static/img/" + filepath.Base(file),
		})
	}

	router := gin.Default()

	router.Static("/static/img", "./static/img")
	router.Static("/site", "./static/site") // TODO make it / without wildcard clashes

	router.GET("/deer", getDeer)

	router.Run("localhost:8080")
}
