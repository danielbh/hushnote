package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	secrets := make(map[string]string)
	secrets["test"] = "test"
	r := gin.Default()
	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true

	r.LoadHTMLGlob("html/*")

	r.GET("/note/:hash", func(c *gin.Context) {
		hash := c.Params.ByName("hash")
		if val, ok := secrets[hash]; ok {
			c.JSON(http.StatusOK, gin.H{"hash": val})
			delete(secrets, hash)
		} else {
			c.String(http.StatusNotFound, "Not Found")
		}
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{})
	})

	r.POST("/submit", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/share")
	})

	r.GET("/share", func(c *gin.Context) {
		c.HTML(http.StatusOK, "share.tmpl", map[string]interface{}{})
	})

	// TODO Add POST / to receive and save secret
	// TODO Add GET / to display shareable link

	r.Run("localhost:8080")
}
