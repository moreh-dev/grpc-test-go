package keycloak

import (
	"context"
	"sync"

	"github.com/Nerzal/gocloak/v13"
)

type KeycloakClient struct {
	Client *gocloak.GoCloak
}

const (
	client_id     = "sdam"
	client_secret = "uBaRfIOFlKlqy4YvXYWCj1Vm0aw0WkI4"
	realm_name    = "moreh"
	username      = "jaeyeon-test"
	userpass      = "moreh@DEV"
)

var (
	client *KeycloakClient
	once   sync.Once
)

func GetKeycloakClient() *KeycloakClient {
	once.Do(func() {
		client = &KeycloakClient{}
		client.Client = gocloak.NewClient("https://keycloak.moreh.dev")
	})

	return client
}

func (kc *KeycloakClient) IssueJWTToken() (*gocloak.JWT, error) {
	ctx := context.Background()

	token, err := kc.Client.Login(ctx, client_id, client_secret, realm_name, username, userpass)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (kc *KeycloakClient) RefreshJWTToken(refreshToken string) (*gocloak.JWT, error) {
	ctx := context.Background()

	token, err := kc.Client.RefreshToken(ctx, refreshToken, client_id, client_secret, realm_name)
	if err != nil {
		return nil, err
	}
	return token, nil
}
