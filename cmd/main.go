package main

import (
	"flag"
	"fmt"
	"regexp"

	"github.com/fprofit/urlMinimizationService/tree/development/pkg/dbpstsql"
	"github.com/fprofit/urlMinimizationService/tree/development/pkg/miniurlgenerator"
	"github.com/gin-gonic/gin"
)

var (
	inMemory = make(map[string]string)
	validUrl = regexp.MustCompile(`^(https?:\/\/)?[\w]{1,32}\.[a-zA-Z]{2,32}[^\s]*$`)
	rGet     = "/getMinimizedUrlToOriginalUrl"
	rPost    = "/postOriginalUrlToMinimizedUrl"
	PsqlInfo = "host=db port=5432 user=postgres password=1234"
)

func main() {
	inMem := flag.Bool("inMemory", false, "Run microService in memory")
	conBDbSQL := flag.Bool("dbSQL", false, "Run DB SQL -dbSQL")
	flag.Parse()

	if *inMem {
		inMemoryFunc()
	} else if *conBDbSQL {
		if dbpstsql.OpenDB(PsqlInfo) {
			funcBDSQL()
		}
	} else {
		fmt.Println("Commands:\n-inMemory\tRun microService in memory\n-dbSQL\tRun DB SQL -dbSQL") // \"host=localhost port=5432 user=postgres password=1234\" '")
	}
}

func funcBDSQL() {

	r := gin.Default()
	r.GET(rGet, func(c *gin.Context) {
		minimizedUrl := c.Query("minimizedUrl")
		originalUrl := dbpstsql.GetOrigToMini(minimizedUrl)

		if originalUrl != "" && originalUrl != "false" {
			c.String(200, originalUrl)
		} else {
			c.String(400, "")
		}
	})

	r.POST(rPost, func(c *gin.Context) {
		originalUrl := c.Query("originalUrl")

		if validUrl.MatchString(originalUrl) {
			miniUrl := dbpstsql.GetMiniToOrig(originalUrl)
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
				miniUrlMap = miniurlgenerator.UrlGenerator(originalUrl)
				inMemory[miniUrlMap] = originalUrl
			}
			c.String(200, miniUrlMap)

		} else {
			c.String(400, "")
		}
	})

	r.Run()
}
