package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello,World!")
	})
	route.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	route.POST("/", func(c *gin.Context) {
		var user User
		buf := make([]byte, 2048)
		n, _ := c.Request.Body.Read(buf)
		if err := json.Unmarshal(buf[0:n], &user); err != nil {
			c.String(http.StatusBadRequest, "json parameter is invalid.")
		} else {
			c.String(http.StatusOK, "Hello!!%s", user.Name)
		}
	})
	route.Run()
}
