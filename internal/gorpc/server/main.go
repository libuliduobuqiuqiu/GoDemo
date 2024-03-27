package main

import (
	"context"
	"fmt"
	"godemo/internal/gorpc/gen"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	gen.UnimplementedSayHelloServer
}

func (g *GrpcServer) Hello(c context.Context, req *gen.HelloReq) (*gen.HelloResp, error) {
	log.Printf("Received rpc requests: ", req.String())
	return &gen.HelloResp{Msg: fmt.Sprintf("hello, world %s", req.Name)}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	gen.RegisterSayHelloServer(server, &GrpcServer{})

	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
