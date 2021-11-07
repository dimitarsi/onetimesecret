package main

import (
	"flag"
	"fmt"

	"github.com/dimitarsi/onetimesecret/api"
	"github.com/dimitarsi/onetimesecret/request"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func main() {

	port := flag.Int("port", 5000, "API Port to listen on")

	flag.Parse()

	app := gin.New()

	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	app.Use(func (c *gin.Context) {

		c.Set("redis", redis.NewClient(&redis.Options{
			Addr: "redis:6379",
			Password: "",
			DB: 0,
		}))

		c.Next()
	})

	app.POST("/create-secret", createSecret)
	app.POST("/secrets", findSecret)

	app.Run(fmt.Sprintf(":%d", *port))
}


func createSecret(c *gin.Context) {
	client, _ := c.Get("redis")

	data := &request.CreateSecretRequest{}

	err := c.BindJSON(data)

	data.Redis = client.(redis.Client)


	if err != nil {
		c.JSON(400, getErrorResponseMessage(err))
		return
	}

	response, err := api.CreateSecret(data)

	if err != nil {
		c.JSON(200, response)
	}

	c.JSON(400, getErrorResponseMessage(err))
}

func findSecret(c *gin.Context) {
	client, _ := c.Get("redis")
	data := &request.FindSecretRequest{}

	err := c.BindJSON(data)

	data.Redis = client.(redis.Client)


	if err != nil {
		c.JSON(400, getErrorResponseMessage(err))
		return 
	}

	response, err := api.FindSecret(data)

	if err != nil {
		c.JSON(400, getErrorResponseMessage(err))
	}

	c.JSON(200, response)
}

func getErrorResponseMessage(err error) map[string]string {
	return map[string]string {
		"status": "error",
		"error": "Something went wrong",
		"details": fmt.Sprintf("%v", err),
	}
}