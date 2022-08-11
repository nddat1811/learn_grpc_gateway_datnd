package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"time"

	demo "gateway/demo"

	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type server struct {
	demo.UnimplementedDemoGatewayServer
}

var jwtKey = os.Getenv("API_KEY")
const (
	expRefreshToken = 24
	expToken        = 15
	company         = "Hybrid Technologies Viet Nam"
	hostMail        = "humghuy201280@gmail.com"
	subjectMail     = "Email reset password"
	textContent     = "Struction to reset your password:"
)


func (server) Register(ctx context.Context, in *demo.RegisterRequest) (*demo.RegisterResponse, error) {
	log.Printf("registering %s\n", in.GetUsername())
	log.Printf("registering %s\n", in.GetPassword())
	username := in.GetUsername()
	password := in.GetPassword()

	if username == "" || password == "" {
		return nil, status.Error(codes.DataLoss, "Lost input") //error grpc.Error(fmt.Sprint)
	}
	// token, err :=
	response := &demo.RegisterResponse{
		Msg: "sspss",
	}
	return response, nil
}

func (m* server) Login(ctx context.Context, in *demo.LoginRequest) (*demo.LoginResponse, error) {
	username := in.GetUsername()
	password := in.GetPassword()

	if username == "" || password == "" {
		return nil, status.Error(codes.DataLoss, "Lost input") //error grpc.Error(fmt.Sprint)
	}

	tokenJwt := jwt.New(jwt.SigningMethodHS256)
	claims := tokenJwt.Claims.(jwt.MapClaims)
	claims["id"] = "1111"
	claims["exp"] = time.Now().Add(time.Hour * expToken).Unix()
	token, err := tokenJwt.SignedString([]byte(jwtKey))
	if err != nil {
		return nil, err
	}


	response := &demo.LoginResponse{
		Msg: token,
	}
	return response, nil
}

//https://fale.io/blog/2021/07/28/cors-headers-with-grpc-gateway
func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

//https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/
func main() {

	// err := gotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	// connDB := config.InitMysql()
	// defer config.CloseConnectDB(connDB)

	// GRPC Server
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	demo.RegisterDemoGatewayServer(s, &server{})
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8089")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8089",
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
		Handler: cors(gwmux),
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())

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
