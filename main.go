package main

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	//"sync/atomic"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/getMinimizedUrlToOriginalUrl", func(c *gin.Context) {
		minimizedUrl := c.Query("minimizedUrl")
		originalUrl := checkUrl(minimizedUrl)
		if originalUrl != "false" {
			c.String(200, func(originalUrl string) (minimizedUrl string) {
				minimizedUrl = originalUrl
				return
			}(UrlGenerator(originalUrl)))
		}
	})

	r.POST("/postOriginalUrlToMinimizedUrl", func(c *gin.Context) {
		originalUrl := c.Query("originalUrl")
		minimizedUrl := checkUrl(originalUrl)
		if minimizedUrl != "false" {
			c.String(200, func(minimizedUrl string) (originalUrl string) {
				originalUrl = minimizedUrl
				return
			}(UrlGenerator(minimizedUrl)))
		}
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func UrlGenerator(originalUrl string) (minimizedUrl string) {

	// if originalUrl == urlToDB{
	// 	minimizedUrl = urlToDB[originalUrl]
	// 	return
	// }

	symbol := "1234567890_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 10; i++ {
		time.Sleep(42 * time.Millisecond)
		rand.Seed(time.Now().UTC().UnixMicro())
		minimizedUrl += string(symbol[rand.Intn(62)])
	}
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
