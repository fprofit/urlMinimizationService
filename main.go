package main

import (
	"math/rand"
	"regexp"
	"time"

	//"database/sql"

	"github.com/gin-gonic/gin"
)

var (
	inMemory = make(map[string]string)
	validUrl = regexp.MustCompile(`^(https?:\/\/)?[\w]{1,32}\.[a-zA-Z]{2,32}[^\s]*$`)
)

func main() {
	r := gin.Default()
	r.GET("/getMinimizedUrlToOriginalUrl", func(c *gin.Context) {
		minimizedUrl := c.Query("minimizedUrl")
		originalUrl := ""
		for origUrl, miniUrl := range inMemory {
			if minimizedUrl == miniUrl {
				originalUrl = origUrl
			}
		}
		if originalUrl != "" {
			c.String(200, originalUrl)
		} else {
			c.String(400, originalUrl)
		}
	})

	r.POST("/postOriginalUrlToMinimizedUrl", func(c *gin.Context) {
		originalUrl := c.Query("originalUrl")
		if validUrl.MatchString(originalUrl) {
			c.String(200, UrlGenerator(originalUrl))
		} else {
			c.String(400, "")
		}
	})

	r.Run()
}

func UrlGenerator(originalUrl string) (minimizedUrl string) {

	_, ok := inMemory[originalUrl]
	if ok {
		return inMemory[originalUrl]
	}

	symbol := "1234567890_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 10; i++ {
		time.Sleep(42 * time.Millisecond)
		rand.Seed(time.Now().UTC().UnixNano())
		minimizedUrl += string(symbol[rand.Intn(62)])
	}

	inMemory[originalUrl] = minimizedUrl

	return
}
