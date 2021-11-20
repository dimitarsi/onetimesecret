package api

import (
	"encoding/json"
	"time"
	"strings"
	"fmt"

	"github.com/dimitarsi/onetimesecret/request"
	"github.com/dimitarsi/onetimesecret/request/validation"
	"github.com/gin-gonic/gin"
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

	k := request.Identity.NewId()
	expires := time.Now().Add(Expire)

	hasErrors, validationErrors := validation.CheckPassword(request.Password)

	if hasErrors {
		return gin.H{}, fmt.Errorf("%s", strings.Join( []string(validationErrors), ","))
	}

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

	err = request.Secrets.Set(k, string(redisVal))

	return gin.H{
		"entry": k,
		"expires": expires,
	}, err
}