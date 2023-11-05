package main

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	//go:embed static/*
	f embed.FS
)

func main() {
	r := gin.Default()
	//	r.GET("/solve", func(c *gin.Context) {
	//		//query := c.Request.URL.Query()
	//		fmt.Println(c.Params)
	//
	//		c.HTML(200, "test", nil)
	//	})
	r.StaticFS("/", http.FS(f))
	r.Run(":8080")
}
