package interceptor

import (
	"context"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"moreh.io/go-grpc-test/internal/pkg/utils"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {

	panicked := true
	defer func() {
		if r := recover(); r != nil || panicked {

		}
	}()

	fmt.Println("auth requested")

	//Do Auth
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("no metadata")
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	tokenString := md.Get("authorization")
	if len(tokenString) == 0 {
		fmt.Println("no auth token")
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	rawToken := strings.TrimPrefix(tokenString[0], "Bearer: ")

	token, err := utils.VerifyJWTToken(rawToken)
	if err != nil {
		fmt.Println("invalid token")
		return nil, status.Errorf(codes.PermissionDenied, "invalid token: %v", err)
	}

	_, ok = token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("invalid token claim")
		return nil, status.Errorf(codes.PermissionDenied, "invalid token claims")
	}
	fmt.Println("token ok")

	return handler(ctx, req)
}
