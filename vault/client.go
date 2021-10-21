package vault

import (
	"fmt"

	"github.com/hashicorp/vault/api"
)

type VaultClient interface {
	GetSecretsV2(path string, args ...interface{}) (map[string]interface{}, error)
}

type vaultClientImpl struct {
	client *api.Client
}

func NewVaultClient(url string, auth Authenticator) (VaultClient, error) {
	client, err := api.NewClient(&api.Config{
		Address: url,
	})
	if err != nil {
		return nil, err
	}

	if auth != nil {
		if err := auth.Authenticate(client); err != nil {
			return nil, err
		}
	}

	v := vaultClientImpl{client: client}
	return v, nil
}

func (v vaultClientImpl) GetSecretsV2(path string, args ...interface{}) (map[string]interface{}, error) {
	secret, err := v.client.Logical().Read(fmt.Sprintf(path, args...))
	if err != nil {
		return nil, err
	}
	m, _ := secret.Data["data"].(map[string]interface{})
	return m, nil
}
