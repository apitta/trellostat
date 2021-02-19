package trellostat

import (
	"errors"
)

// TrelloCredential is the credential to connect to the API
type TrelloCredential struct {
	Token string
	Key   string
}

// TrelloCredentials interact with user and collects key and token
func TrelloCredentials(key, token string) (TrelloCredential, error) {
	if len(key) == 0 {
		return TrelloCredential{}, errors.New("Key is required")
	}
	if len(token) == 0 {
		return TrelloCredential{}, errors.New("Token is required")
	}
	credential := TrelloCredential{
		Token: token,
		Key:   key,
	}
	return credential, nil
}
