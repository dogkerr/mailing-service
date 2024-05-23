package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/dogkerr/mailing-service/m/v2/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddress := flag.String("server", "localhost:8080", "server address")
	flag.Parse()
	log.Printf("start client on server %s", *serverAddress)

	conn, err := grpc.NewClient(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect to server: %v", err)
	}

	emailClient := pb.NewEmailServiceClient(conn)

	req := &pb.VerificationEmailRequest{
		Email:           "davidlou0810@gmail.com",
		Name:            "David Lou",
		VerifictionLink: "https://example.com",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := emailClient.SendVerificationEmail(ctx, req)
	if err != nil {
		log.Fatalf("cannot send email: %v", err)
	}

	log.Printf(res.Message)
}
