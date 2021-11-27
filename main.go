package main

import (
	//"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/getMinimizedUrlToOriginalUrl", func(c *gin.Context) {
		c.String(200, func(minimizedUrl string) (originalUrl string) {
			originalUrl = minimizedUrl
			return
		}("get"))
	})

	r.POST("/postOriginalUrlToMinimizedUrl", func(c *gin.Context) {
		c.String(200, func(originalUrl string) (minimizedUrl string) {
			minimizedUrl = originalUrl
			return
		}("post"))
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
		rand.Seed(time.Now().UTC().UnixNano() % 10000)
		minimizedUrl += string(symbol[rand.Intn(62)])
	}
	return
}
