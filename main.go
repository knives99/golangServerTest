package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	limiter "github.com/julianshen/gin-limiter"
	"server_test/Controller"
	"server_test/Tool"
	"time"
)

func main() {
	cfg, err := Tool.ParseConfig("./Config/app.json")
	if err != nil {
		panic(err.Error())
	}
	engine, lm := gin.Default(), limiter.NewRateLimiter(time.Minute, 10, func(ctx *gin.Context) (string, error) {
		key := ctx.Request.Header.Get("X-API-KEY")
		if key == "" {
			return key, nil
		}
		return "", errors.New("API key is missing")
	})

	v1 := engine.Group("/v1")
	v1.Use(lm.Middleware())

	{
		v1.GET("/hello", Controller.GetTopic)
	}

	engine.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//func RateLimiter(interval time.Duration, capacity int64, keyGen limiter.RateKeyFunc) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//
//	}
//}
