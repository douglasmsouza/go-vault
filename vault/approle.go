package vault

import "github.com/hashicorp/vault/api"

type AppRoleAuth struct {
	RoleID   string
	SecretID string
}

func (a AppRoleAuth) Authenticate(client *api.Client) error {
	resp, err := client.Logical().Write("auth/approle/login", map[string]interface{}{
		"role_id":   a.RoleID,
		"secret_id": a.SecretID,
	})
	if err != nil {
		return  err
	}

	client.SetToken(resp.Auth.ClientToken)
	return nil
}
