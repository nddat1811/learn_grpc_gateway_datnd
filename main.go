package main

import (
	"context"
	"fmt"
	"log"
	"net"

	demo "gateway/demo"

	"google.golang.org/grpc"
)

type server struct {
	demo.UnimplementedDemoGatewayServer
}

func (server) Echo(ctx context.Context, msg *demo.StringMessage) (*demo.StringMessage, error) {
	log.Printf("receive message %s\n", msg.GetMsg())
	//msg = msg.GetMsg() + "soss"
	return msg, nil
}

func (server) Register(ctx context.Context, in *demo.RegisterRequest) (*demo.RegisterResponse, error) {
	response := &demo.RegisterResponse{
		Msg: "sspss",
	}
	return response, nil
}

func (server) Login(ctx context.Context, in *demo.LoginRequest) (*demo.LoginResponse, error) {
	response := &demo.LoginResponse{
		Msg: "login ne",
	}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:4002")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	demo.RegisterDemoGatewayServer(s, &server{})

	fmt.Println("demo gateway service is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
