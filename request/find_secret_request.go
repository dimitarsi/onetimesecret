package request

type FindSecretRequest struct {
	GenericSecretRequest

	SecretId string
	Hash     string
}