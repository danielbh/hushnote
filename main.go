package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	secrets := make(map[string]string)

	secrets["test"] = "test"

	r := gin.Default()
	r.GET("/:hash", func(c *gin.Context) {
		hash := c.Params.ByName("hash")

		if val, ok := secrets[hash]; ok {
			c.JSON(http.StatusOK, gin.H{"hash": val})
			delete(secrets, hash)
		} else {
			c.String(http.StatusNotFound, "Not Found")
		}
	})

	// TODO Add GET / for form to enter secret
	// TODO Add POST / to receive and save secret

	r.Run("localhost:8080")
}
