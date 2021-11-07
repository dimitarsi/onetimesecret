package api

import (
	"encoding/json"
	"fmt"

	"github.com/dimitarsi/onetimesecret/request"
	"golang.org/x/crypto/bcrypt"
)

// FindSecret endpoint accepts the following parameters:
//
// @param string SecretId - the uuid generated from createSecret
// @param string Hash - plain text password, used to verify the user
//
// @see request.findSecretRequest
func FindSecret(request *request.FindSecretRequest) (map[string]string, error) {

	rawData, err := request.Redis.Get(request.SecretId).Result()

	if err != nil {
		fmt.Printf("Error finding redis key - %s", request.SecretId)
		return nil, err
	}

	data := make(map[string]string)
	
	err = json.Unmarshal([]byte(rawData), &data)


	if err != nil {
		fmt.Printf("Error unmarshalling the redis data")
		return nil, fmt.Errorf("no such secret")
	}

	err = bcrypt.CompareHashAndPassword([]byte(data["password"]), []byte(request.Hash))
	
	if err != nil {
		fmt.Printf("Passwords didn't match, data[\"password\"]=%d len; request[\"password\"]=%d", len(data["password"]), len(request.Hash))
		return nil, fmt.Errorf("no such secret")
	}

	request.Redis.Del(request.SecretId)

	return map[string]string{
		"message": data["message"],
	}, nil
	
}