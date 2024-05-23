package repository

import (
	"context"
	"log"
	"time"

	"github.com/dogkerr/mailing-service/m/v2/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type service struct{}

type Service interface {
	GetUserById(userID string) (*pb.User, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetUserById(userID string) (*pb.User, error) {
	//users gRPC Client Setup
	conn, err := grpc.NewClient("10.66.66.1:4001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	usersClient := pb.NewUsersServiceClient(conn)

	req := &pb.GetUserRequest{
		Id: userID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Get User by ID
	user, err := usersClient.GetUserById(ctx, req)

	if err != nil {
		return nil, err
	}

	return user, nil
}
