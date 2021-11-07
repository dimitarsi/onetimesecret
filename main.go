package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/dimitarsi/onetimesecret/request"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

func main() {

	port := flag.Int("port", 5000, "API Port to listen on")

	flag.Parse()

	app := gin.New()

	app.POST("/create-secret", createSecret)

	app.Run(fmt.Sprintf(":%d", *port))
}

const (
	Expire time.Duration = time.Minute * 20
)


func createSecret(c *gin.Context) {
	data := &request.CreateSecretRequest{}
	err := c.BindJSON(data)

	if err != nil {
		c.JSON(400, getErrorResponseMessage(err))
		return
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})

	k, _ := uuid.NewUUID()

	redisClient.Set(k.String(), data, Expire)

	expires := time.Now().Add(Expire)

	c.JSON(200, gin.H{
		"secret": k.String(),
		"expires": expires.Local().String(),
	})
}

func getErrorResponseMessage(err error) map[string]string {
	return map[string]string {
		"status": "error",
		"error": "Something went wrong",
		"details": fmt.Sprintf("%v", err),
	}
}