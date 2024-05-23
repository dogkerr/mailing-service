package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/dogkerr/mailing-service/m/v2/pb"
	"github.com/dogkerr/mailing-service/m/v2/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "port to connect to")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	emailServer := service.NewEmailServer()
	grpcServer := grpc.NewServer()
	pb.RegisterEmailServiceServer(grpcServer, emailServer)

	address := fmt.Sprintf("localhost:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
