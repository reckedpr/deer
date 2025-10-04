package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var deerList []Deer
var factList []Fact

type Deer struct {
	ImgURL string `json:"img_url,omitempty"`
}

type Fact struct {
	Text string `json:"fact"`
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

func getFact(c *gin.Context) {
	randomIndex := rand.Intn(len(factList))
	fact := factList[randomIndex]

	c.IndentedJSON(http.StatusOK, fact)
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

	for _, file := range files {
		deerList = append(deerList, Deer{
			ImgURL: "/static/img/" + filepath.Base(file),
		})
	}

	factFile, err := os.ReadFile("data/facts.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(factFile, &factList)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.Static("/static/img", "./static/img")

	router.GET("/", getDeer)
	router.GET("/fact", getFact)

	router.Run("localhost:8080")
}
