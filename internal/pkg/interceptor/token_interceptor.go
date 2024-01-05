package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const token = ""

func AddTokenInterceptor(ctx context.Context, method string, req interface{}, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	newCtx := metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer: "+token)

	fmt.Println("Added token")

	return invoker(newCtx, method, req, reply, cc, opts...)
}
