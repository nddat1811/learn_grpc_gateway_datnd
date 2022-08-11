package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/golang/glog"
	//"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"

	gw "gateway/demo"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:4002", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{}
	err := gw.RegisterDemoGatewayHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8083", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	log.Println("proxy is running...")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{}
	err := gw.RegisterDemoGatewayHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		glog.Fatal(err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.ListenAndServe(":8083", mux)

}