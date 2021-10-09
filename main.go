package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	template, _ := filepath.Abs("templates/index.html")
	r := gin.Default()
	r.LoadHTMLFiles(template)

	r.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.Run(":80")
}
