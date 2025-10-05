package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const (
	imgPath  = "./static/img"
	factPath = "./data/facts.json"
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

func respondError(c *gin.Context, code int, msg string, err error) {
	if err != nil {
		log.Printf("[ %d ] %s: %v", code, msg, err)
	}
	c.JSON(code, gin.H{"error": msg})
	c.Abort()
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
	err := readFacts(factPath)
	if err != nil {
		respondError(c, http.StatusInternalServerError, "failed to get random fact", err)
		return
	}

	randomIndex := rand.Intn(len(factList))
	fact := factList[randomIndex]

	c.IndentedJSON(http.StatusOK, fact)
}

func readFacts(factPath string) error {
	factFile, err := os.ReadFile(factPath)
	if err != nil {
		return fmt.Errorf("failed to read fact path @ %s\n\n%s", factPath, err)
	}

	err = json.Unmarshal(factFile, &factList)
	if err != nil {
		return fmt.Errorf("failed to open %s\n%s", factFile, err)
	}

	return nil
}

func main() {
	var (
		err   error
		files []string
	)

	files, err = FilePathWalkDir(imgPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		deerList = append(deerList, Deer{
			ImgURL: "/img/" + filepath.Base(file),
		})
	}

	readFacts(factPath)

	router := gin.Default()

	router.Static("/img", imgPath)

	router.GET("/", getDeer)
	router.GET("/fact", getFact)

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./web/public/favicon.ico")
	})

	router.Run("localhost:8080")
}
