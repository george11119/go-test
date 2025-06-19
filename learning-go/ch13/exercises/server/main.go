package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var PORT = ":3001"

func JSONLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			log := make(map[string]interface{})

			log["status_code"] = params.StatusCode
			log["path"] = params.Path
			log["method"] = params.Method
			log["start_time"] = params.TimeStamp.Format("2006/01/02 - 15:04:05")
			log["ip"] = params.ClientIP
			log["response_time"] = params.Latency.String()

			s, _ := json.Marshal(log)
			return string(s) + "\n"
		},
	)
}

func main() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(JSONLoggerMiddleware())

	router.GET("/", func(c *gin.Context) {
		t := time.Now()

		if c.Request.Header["Accept"][0] == "application/json" {
			c.JSON(http.StatusOK, gin.H{
				"day_of_week":  t.Weekday().String(),
				"day_of_month": t.Day(),
				"month":        t.Month().String(),
				"year":         t.Year(),
				"hour":         t.Hour(),
				"minute":       t.Minute(),
				"second":       t.Second(),
			})
		} else {
			c.String(http.StatusOK, t.Format(time.RFC3339))
		}
	})

	router.Run(PORT)
}
