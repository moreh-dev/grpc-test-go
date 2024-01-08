package interceptor

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"moreh.io/go-grpc-test/internal/pkg/keycloak"
)

var token *gocloak.JWT

func AddTokenInterceptor(ctx context.Context, method string, req interface{}, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	keycloakClient := keycloak.GetKeycloakClient()
	if token == nil {
		newToken, err := keycloakClient.IssueJWTToken()
		if err != nil {
			return err
		}
		token = newToken
	} else if isTokenExpired() {
		newToken, err := keycloakClient.RefreshJWTToken(token.RefreshToken)
		if err != nil {
			return err
		}
		token = newToken
	}

	newCtx := metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer: "+token.AccessToken)
	fmt.Println("Added token")

	return invoker(newCtx, method, req, reply, cc, opts...)
}

func isTokenExpired() bool {
	if token == nil {
		return true
	}
	//TODO: add token expiration logic

	return false
}
