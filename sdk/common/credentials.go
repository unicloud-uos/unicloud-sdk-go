package common

type Credential struct {
	SecretId  string
	SecretKey string
	Token     string
}

func NewCredential(secretId, secretKey string) *Credential {
	return &Credential{
		SecretId:  secretId,
		SecretKey: secretKey,
	}
}

