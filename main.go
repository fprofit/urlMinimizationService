package main

import (
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	inMemory = make(map[string]string)
	validUrl = regexp.MustCompile(`^(https?:\/\/)?[\w]{1,32}\.[a-zA-Z]{2,32}[^\s]*$`)
	rGet     = "/getMinimizedUrlToOriginalUrl"
	rPost    = "/postOriginalUrlToMinimizedUrl"
)

func main() {
	inMem := flag.Bool("inMemory", false, "Run microService in memory")
	conBDbSQL := flag.String("dbSQL", "", "Run DB SQL")
	flag.Parse()

	if *inMem {
		fmt.Println("inMemory", *inMem)
		inMemoryFunc()
	} else if *conBDbSQL != "" {
		psqlInfo = *conBDbSQL
		if openDB() {
			funcBDSQL()
		}
	}
}

func funcBDSQL() {

	r := gin.Default()
	r.GET(rGet, func(c *gin.Context) {
		minimizedUrl := c.Query("minimizedUrl")
		originalUrl := getOrigToMini(minimizedUrl)

		if originalUrl != "" && originalUrl != "false" {
			c.String(200, originalUrl)
		} else {
			c.String(400, originalUrl)
		}
	})

	r.POST(rPost, func(c *gin.Context) {
		originalUrl := c.Query("originalUrl")

		if validUrl.MatchString(originalUrl) {
			miniUrl := getMiniToOrig(originalUrl)
			if miniUrl != "false" {
				c.String(200, miniUrl)
			} else {
				c.String(400, "")
			}

		} else {
			c.String(400, "")
		}
	})

	r.Run()
}

func inMemoryFunc() {

	r := gin.Default()
	r.GET(rGet, func(c *gin.Context) {
		minimizedUrl := c.Query("minimizedUrl")
		originalUrl := ""
		_, ok := inMemory[minimizedUrl]
		if ok {
			originalUrl = inMemory[minimizedUrl]
			c.String(200, originalUrl)
		} else {
			c.String(400, "")
		}
	})

	r.POST(rPost, func(c *gin.Context) {
		originalUrl := c.Query("originalUrl")
		if validUrl.MatchString(originalUrl) {
			miniUrlMap := ""
			for miniUrl, origUrl := range inMemory {
				if origUrl == originalUrl {
					miniUrlMap = miniUrl
				}
			}
			if miniUrlMap == "" {
				miniUrlMap = UrlGenerator(originalUrl)
				inMemory[miniUrlMap] = originalUrl
			}
			c.String(200, miniUrlMap)

		} else {
			c.String(400, "")
		}
	})

	r.Run()
}

func UrlGenerator(originalUrl string) (minimizedUrl string) {
	symbol := "1234567890_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 10; i++ {
		time.Sleep(42 * time.Millisecond)
		rand.Seed(time.Now().UTC().UnixNano())
		minimizedUrl += string(symbol[rand.Intn(62)])
	}
	return
}
