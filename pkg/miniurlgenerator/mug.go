package miniurlgenerator

import (
	"math/rand"
	"time"
)

func UrlGenerator(originalUrl string) (minimizedUrl string) {
	symbol := "1234567890_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 10; i++ {
		time.Sleep(42 * time.Millisecond)
		rand.Seed(time.Now().UTC().UnixNano())
		minimizedUrl += string(symbol[rand.Intn(62)])
	}
	return
}
