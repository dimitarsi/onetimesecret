package api

import (
	"encoding/json"
	"testing"

	"github.com/dimitarsi/onetimesecret/repository"
	"github.com/dimitarsi/onetimesecret/request"
	"github.com/dimitarsi/onetimesecret/utils"
	"github.com/golang/mock/gomock"
)

func TestCreateSecret(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	secretsMock := repository.NewMockSecretRepository(ctrl)
	identityMock := utils.NewMockIdentityUtil(ctrl)

	mockId := "uuid-1234"
	mockPassword := "Foobar"
	mockMessage := "User"

	identityMock.EXPECT().NewId().Return(mockId)

	jsonValue,_ := json.Marshal(map[string]string{
		"message": mockMessage,
		"password":  mockPassword,
	})

	// Ensures we are not saving the password in plain text
	secretsMock.EXPECT().Set(mockId, gomock.Not(jsonValue)).Times(1)
	secretsMock.EXPECT().GetDel(mockId).Times(0)

	request := &request.CreateSecretRequest{
		Message: mockMessage,
		Password: mockPassword,
	}

	request.Secrets = secretsMock
	request.Identity = identityMock

	CreateSecret(request)
}