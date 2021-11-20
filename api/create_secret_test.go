package api

import (
	"encoding/json"
	"testing"

	"github.com/dimitarsi/onetimesecret/repository"
	"github.com/dimitarsi/onetimesecret/request"
	"github.com/dimitarsi/onetimesecret/utils"
	"github.com/golang/mock/gomock"
)

func TestCreateSecretWithWeakPassword(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	secretsMock := repository.NewMockSecretRepository(ctrl)
	identityMock := utils.NewMockIdentityUtil(ctrl)

	mockPassword := "Foobar"
	mockMessage := "User"

	request := &request.CreateSecretRequest{
		Message: mockMessage,
		Password: mockPassword,
	}

	request.Secrets = secretsMock
	request.Identity = identityMock

	_, err := CreateSecret(request)

	if err == nil {
		t.Fatal("CreateSecret should return an error when weak password is used")
	}
}

func TestCreateSecretWithStrongPassword(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	secretsMock := repository.NewMockSecretRepository(ctrl)
	identityMock := utils.NewMockIdentityUtil(ctrl)

	mockId := "uuid-1234"
	mockPassword := "ThisIsALongPassword!"
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