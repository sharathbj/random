package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"random/golang/grpc/protos"
)

func main()  {
	fmt.Println("client started")
	conn, err := grpc.Dial(":7778",grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to dial",err)
	}
	defer conn.Close()
	fmt.Println("client up and running :)")
	client := protos.NewPingServiceClient(conn)
	req := &protos.PingStruct{Req:"hello world"}
	resp ,err:= client.PingCheck(context.Background(),req)
	fmt.Println(err,resp.Resp)
}