package main

import (
	"github.com/gin-gonic/gin"

	"newsfeeder/httpd/handler"
	"newsfeeder/platform/newsfeed"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/pinga", handler.PingGet())

	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
