package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"moreh.io/go-grpc-test/internal/pkg/interceptor"
	"moreh.io/go-grpc-test/proto"
)

func main() {
	serverAddr := "localhost:25050"
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor.AddTokenInterceptor))
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	defer conn.Close()

	sdamAPIClient := proto.NewUserAPIClient(conn)

	data, err := sdamAPIClient.ListSDAModel(context.Background(), &proto.Token{Value: "dummy"})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
}
