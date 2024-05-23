package service

import (
	"context"
	"fmt"
	"log"

	"github.com/dogkerr/mailing-service/m/v2/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EmailServer struct {
	pb.UnimplementedEmailServiceServer
}

func NewEmailServer() *EmailServer {
	return &EmailServer{}
}

func (server *EmailServer) SendBillingEmail(ctx context.Context, req *pb.BillingEmailRequest) (*pb.BillingEmailResponse, error) {
	user := req.GetUserId()
	log.Printf("received request for user %s", user)

	if ctx.Err() == context.Canceled {
		log.Printf("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}

	res := &pb.BillingEmailResponse{
		Message: fmt.Sprintf("email successfully sent to %s", user),
	}
	return res, nil
}

func (server *EmailServer) SendVerificationEmail(ctx context.Context, req *pb.VerificationEmailRequest) (*pb.VerificationEmailResponse, error) {
	email := req.GetEmail()
	log.Printf("received request for email %s", email)

	if ctx.Err() == context.Canceled {
		log.Printf("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}

	res := &pb.VerificationEmailResponse{
		Message: fmt.Sprintf("email successfully sent to %s", email),
	}
	return res, nil
}
