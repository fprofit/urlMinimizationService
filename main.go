package main

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	//"sync/atomic"

	"github.com/gin-gonic/gin"
)

var inMemory = make(map[string]string)

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
		if checkUrl(originalUrl) != "false" {
			c.String(200, UrlGenerator(checkUrl(originalUrl)))
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
		rand.Seed(time.Now().UTC().UnixMicro())
		minimizedUrl += string(symbol[rand.Intn(62)])
	}

	inMemory[originalUrl] = minimizedUrl

	return
}

func checkUrl(url string) string {
	var validUrl = regexp.MustCompile(`(https?:\/\/)?[\w-]{1,32}\.[a-zA-Z]{2,32}[^\s]*`)
	var validHttp = regexp.MustCompile(`(https?:\/\/)[\w-]{1,32}\.[a-zA-Z]{2,32}[^\s]*`)
	var validWww = regexp.MustCompile(`(www.)[\w-]{1,32}\.[a-zA-Z]{2,32}[^\s]*`)
	if validUrl.MatchString(url) {
		if validHttp.MatchString(url) {
			strHttp := strings.Split(url, "://")
			url = strHttp[1]
		}
		if validWww.MatchString(url) {
			strWww := strings.Split(url, "www.")
			url = strWww[1]
		}
		return url
	}
	return "false"
}
