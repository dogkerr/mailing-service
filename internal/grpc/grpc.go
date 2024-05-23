package grpc

import (
	"net"

	"github.com/dogkerr/mailing-service/m/v2/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGRPCServer(
	emailServer pb.EmailServiceServer,
	listener net.Listener,
	ch chan *grpc.Server,
) error {
	// GRPC Server
	grpcServer := grpc.NewServer()
	pb.RegisterEmailServiceServer(grpcServer, emailServer)
	reflection.Register(grpcServer)

	ch <- grpcServer

	return grpcServer.Serve(listener)
}
