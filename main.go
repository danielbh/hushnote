package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	secrets := make(map[string]string)
	r := gin.Default()

	r.LoadHTMLGlob("html/*")

	r.GET("/note/:hash", func(c *gin.Context) {
		hash := c.Params.ByName("hash")
		if val, ok := secrets[hash]; ok {
			c.HTML(http.StatusOK, "note.tmpl", gin.H{
				"note": val,
			})
			delete(secrets, hash)
		} else {
			c.String(http.StatusNotFound, "Not Found")
		}
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{})
	})

	r.POST("/share", func(c *gin.Context) {
		note := c.PostForm("note")
		key := uuid.New().String()
		secrets[key] = note

		c.HTML(http.StatusOK, "share.tmpl", gin.H{
			"url": c.Request.Host + "/note/" + key,
		})
	})

	r.Run("localhost:8080")
}
