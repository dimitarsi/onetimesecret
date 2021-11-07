package request

type CreateSecretRequest struct {
	GenericSecretRequest

	Message  string
	Password string
}