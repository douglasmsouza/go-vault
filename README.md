# Go Vault
A simple helper for accessing [Vault API](https://github.com/hashicorp/vault/tree/main/api).

Currently supports authenticate with AppRole and retrieving kv version 2 secrets.

# Usage
Authenticate with approle and retrieving a secret

```go
package main

import (
	"fmt"

	"github.com/douglasmsouza/go-vault/vault"
)

func main() {
	auth := vault.AppRoleAuth{
		RoleID:   "your role id",
		SecretID: "your secret id",
	}
	url := "your vault url"
	client, _ := vault.NewVaultClient(url, auth)
	secrets, _ := client.GetSecretsV2("secrets/data/mysecret")
	fmt.Printf("%v", secrets)
}

```