package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"random/golang/grpc/protos"
)

func main()  {
	fmt.Println("server starup")
	lis, err := net.Listen("tcp", ":7778")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	protos.RegisterPingServiceServer(grpcServer,&protos.SamplePing{})
	fmt.Println("server listening")
	grpcServer.Serve(lis)
}