package vault

import "github.com/hashicorp/vault/api"

type Authenticator interface {
	Authenticate(client *api.Client) error
}

