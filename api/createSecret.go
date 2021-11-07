package api

import (
	"time"

	"github.com/dimitarsi/onetimesecret/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	Expire time.Duration = time.Minute * 20
	HashCost int = 10
)

func CreateSecret(request *request.CreateSecretRequest) (map[string]interface{}, error) {

	k, _ := uuid.NewUUID()
	expires := time.Now().Add(Expire)

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), HashCost)

	if err != nil {
		return gin.H{}, err
	}

	status := request.Redis.Set(k.String(), map[string]string{
		"message": request.Message,
		"password": string(hash),
	}, Expire)

	return gin.H{
		"entry": k,
		"expires": expires,
	}, status.Err()
}