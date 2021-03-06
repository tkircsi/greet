package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tkircsi/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	grpcHost = "tkircsi.net:50052"
)

func main() {
	fmt.Println("Greeting RPC client started....")
	certFile := "ssl/tkircsi_net.crt"
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatalf("error while loading ca trust certificate: %v", sslErr)
	}

	opts := grpc.WithTransportCredentials(creds)
	cc, err := grpc.Dial(grpcHost, opts)
	if err != nil {
		log.Fatalf("error while connecting to grpc server: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doGreet(c)
}

func doGreet(c greetpb.GreetServiceClient) {
	resp, err := c.Greet(context.Background(), &greetpb.GreetRequest{
		FirstName: "János",
		LastName:  "Kelemen",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.GetResponse())
}
