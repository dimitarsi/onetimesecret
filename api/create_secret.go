package api

import (
	"encoding/json"
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

// CreateSecret endpoint accepts the following parameters
// @param string Password
// @param string Message
//
// @see requests.createSecretRequest
func CreateSecret(request *request.CreateSecretRequest) (map[string]interface{}, error) {

	k, _ := uuid.NewUUID()
	expires := time.Now().Add(Expire)

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), HashCost)

	if err != nil {
		return gin.H{}, err
	}

	redisVal, err := json.Marshal(map[string]string{
		"message": request.Message,
		"password": string(hash),
	})

	if err != nil {
		return gin.H{}, err
	}

	status := request.Redis.Set(k.String(), string(redisVal) , Expire)

	return gin.H{
		"entry": k,
		"expires": expires,
	}, status.Err()
}