package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Wiki struct {
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
		// var samples []Config
		buf := make([]byte, 2048)
		n, _ := c.Request.Body.Read(buf)
		var lines []string
		var wikis []Wiki
		var jsonline Wiki
		for i := range lines {
			if lines[i] != "" {
				//if err := json.Unmarshal([]byte(lines[i]), &wikis); err != nil {
				if err := json.Unmarshal([]byte(lines[i]), &jsonline); err != nil {
					log.Fatal(err)
					wikis = append(wikis, jsonline)
				}
			}
		}
		// err := json.Unmarshal(buf[0:n], &samples)
		c.String(http.StatusOK, "Hello!! %s", wikis)
	})
	r.Run()
}
