package main

import (
	"fmt"
	"log"
	"mygin/api"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting API server...")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/something", asJSON(func(c *gin.Context) (interface{}, error) {
		return api.GetSomething(c.Query("project"), c.Query("version"))
	}))

	r.GET("/projects/:project/:version", asJSON(func(c *gin.Context) (interface{}, error) {
		return api.GetSomething(c.Param("project"), c.Param("version"))
	}))

	r.GET("/externals/:project/:version", asJSON(func(c *gin.Context) (interface{}, error) {

		project := c.Param("project")
		if !strings.HasPrefix(project, "ex_") {
			return nil, fmt.Errorf("project needs to start with 'ex_'")
		}
		return api.GetSomething(c.Param("project"), c.Param("version"))
	}))

	r.Run()
}

// asJSON provides a reusable implementation of a JSON end point
func asJSON(handler func(*gin.Context) (interface{}, error)) func(*gin.Context) {
	return func(c *gin.Context) {
		data, err := handler(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if data == nil {
			c.JSON(http.StatusNoContent, gin.H{})
		}

		c.JSON(http.StatusOK, data)
	}
}
