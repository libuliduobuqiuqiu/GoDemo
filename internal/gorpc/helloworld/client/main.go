package main

import (
	"context"
	"godemo/internal/gorpc/helloworld/gen"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	con, err := grpc.Dial("127.0.0.1:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	client := gen.NewSayHelloClient(con)

	resp, err := client.Hello(context.Background(), &gen.HelloReq{Name: "zhangsan"})
	if err != nil {
		log.Fatal(resp)
	}

	log.Println(resp.Msg)
}
