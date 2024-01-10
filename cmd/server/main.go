package main

import (
	"fmt"
	"net"
	"os"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"moreh.io/go-grpc-test/internal/pkg/interceptor"
	sdamanager_api "moreh.io/go-grpc-test/internal/pkg/service"
	moreh_proto "moreh.io/go-grpc-test/proto"
)

func main() {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(interceptor.AuthInterceptor),
		),
	)
	moreh_proto.RegisterUserAPIServer(grpcServer, sdamanager_api.NewSDAManagerUserAPI())
	moreh_proto.RegisterJobAPIServer(grpcServer, sdamanager_api.NewSDAManagerJobAPI())

	errCh := make(chan error)
	managingChan := make(chan os.Signal, 1)

	go func() {
		lis, err := net.Listen("tcp", ":"+"25050")
		fmt.Println("server running on 25050")
		err = grpcServer.Serve(lis)
		if err != nil {
			return
		}

	}()

	for {
		select {
		case sig := <-managingChan:
			switch sig {
			case syscall.SIGUSR1:
			case syscall.SIGUSR2:
			default:
				grpcServer.GracefulStop()
				return
			}
		case err := <-errCh:
			fmt.Println(err)
			return
		}
	}
}
