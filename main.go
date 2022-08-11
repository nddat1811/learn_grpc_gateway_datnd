package main

import (
	"context"
	"log"
	"net"
	"net/http"

	demo "gateway/demo"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	log.Printf("registering %s\n", in.GetUsername())
	log.Printf("registering %s\n", in.GetPassword())
	response := &demo.RegisterResponse{
		Msg: "sspss",
	}
	return response, nil
}

func (server) Login(ctx context.Context, in *demo.LoginRequest) (*demo.LoginResponse, error) {
	log.Printf("login: %s\n", in.GetUsername())
	log.Printf("login: %s\n", in.GetPassword())
	response := &demo.LoginResponse{
		Msg: "login ne",
	}
	return response, nil
}

// func main() {
// 	lis, err := net.Listen("tcp", "0.0.0.0:4002")
// 	if err != nil {
// 		log.Fatalf("err while create listen %v", err)
// 	}

// 	s := grpc.NewServer()

// 	demo.RegisterDemoGatewayServer(s, &server{})

// 	fmt.Println("demo gateway service is running...")
// 	err = s.Serve(lis)

// 	if err != nil {
// 		log.Fatalf("err while serve %v", err)
// 	}
// }

//https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/
func main() {
	// GRPC Server
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	demo.RegisterDemoGatewayServer(s, &server{})
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = demo.RegisterDemoGatewayHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())

}
