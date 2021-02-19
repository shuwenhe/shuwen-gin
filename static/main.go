package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/file", "./dir")
	r.StaticFS("/s", http.Dir("s"))
	r.StaticFile("/f", "./f")
	r.Run()
}
