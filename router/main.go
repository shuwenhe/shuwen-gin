package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(200, "get")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "delete")
	})
	r.Handle("DELETE", "/delete2", func(c *gin.Context) {
		c.String(200, "delete2")
	})
	r.Any("/any", func(c *gin.Context) {
		c.String(200, "any")
	})
	r.Run()
}
