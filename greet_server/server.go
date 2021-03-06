package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/tkircsi/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	grpcPort = ":50051"
	certFile = ""
	keyFile  = ""
)

type server struct{}

func init() {
	if v, ok := os.LookupEnv("GRPC_PORT"); ok {
		grpcPort = v
	}
	if v, ok := os.LookupEnv("CERT_FILE"); ok {
		certFile = v
	}
	if v, ok := os.LookupEnv("KEY_FILE"); ok {
		keyFile = v
	}
}

func main() {
	fmt.Println("Greeting RPC service is started....")

	lis, err := net.Listen("tcp", "0.0.0.0"+grpcPort)
	if err != nil {
		log.Fatalf("invalid port number: %v", err)
	}

	// certFile := "ssl/tkircsi_net.crt"
	// keyFile := "ssl/tkircsi_net.pem"
	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if sslErr != nil {
		log.Fatalf("failed loading certificates: %v", sslErr)
	}

	opts := grpc.Creds(creds)
	s := grpc.NewServer(opts)
	greetpb.RegisterGreetServiceServer(s, &server{})

	log.Fatal(s.Serve(lis))
}

func (*server) Greet(ctx context.Context, r *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet service is called with request: %v\n", r)
	fname := r.GetFirstName()
	if fname == "" {
		fname = "Guest"
	}
	resp := fmt.Sprintf("Greeting %s! Welcome at the GreetService gRPC service.", fname)
	return &greetpb.GreetResponse{
		Response: resp,
	}, nil
}
