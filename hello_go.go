package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello,World!")
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	r.POST("/", func(c *gin.Context) {
		var samples []Config
		buf := make([]byte, 2048)
		n, _ := c.Request.Body.Read(buf)
		json.Unmarshal(buf[0:n], &samples)
		c.String(http.StatusOK, "Hello!! %s", samples)
	})
	r.Run()
}
