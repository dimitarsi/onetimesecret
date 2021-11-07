package api

import (
	"fmt"

	"github.com/dimitarsi/onetimesecret/request"
)

func FindSecret(request *request.FindSecretRequest) (map[string]string, error) {
	data, err := request.Redis.Get(request.SecretId).Result()

	if err == nil {
		return map[string]string{
			"message": data,
		}, nil
	}

	return nil, fmt.Errorf("no such secret")
}