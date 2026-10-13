package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"godemo/internal/gorpc/helloworld/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	con, err := grpc.DialContext(ctx, "127.0.0.1:8090",
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("connect to Hello World server: %w", err)
	}
	defer con.Close()

	client := gen.NewSayHelloClient(con)

	resp, err := client.Hello(ctx, &gen.HelloReq{Name: "zhangsan"})
	if err != nil {
		return fmt.Errorf("call Hello: %w", err)
	}

	log.Println(resp.Msg)
	return nil
}
