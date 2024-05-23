package service

import (
	"context"
	"fmt"
	"log"

	"github.com/dogkerr/mailing-service/m/v2/domain"
	"github.com/dogkerr/mailing-service/m/v2/internal/repository"
	"github.com/dogkerr/mailing-service/m/v2/pb"
	"github.com/dogkerr/mailing-service/m/v2/service/helpers"
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
	log.Printf("received request for user with id %s", req.GetUserId())

	user, err := repository.NewService().GetUserById(req.GetUserId())
	if err != nil {
		log.Printf("error getting user: %v", err)
		return nil, status.Error(codes.Internal, "error getting user")
	}

	templateData := domain.BillingNoticeData{
		Name:                 user.GetFullname(),
		Email:                user.GetEmail(),
		ContainerID:          req.GetContainerId(),
		TotalCPUUsage:        req.GetTotalCpuUsage(),
		TotalMemoryUsage:     req.GetTotalMemoryUsage(),
		TotalNetIngressUsage: req.GetTotalNetIngressUsage(),
		TotalNetEgressUsage:  req.GetTotalNetEgressUsage(),
		Timestamp:            req.GetTimestamp().String(),
		TotalCost:            req.GetTotalCost(),
	}

	err = helpers.ParseAndSend("templates/billing_notice.html", templateData, user.GetEmail(), "Billing Notice Email")
	if err != nil {
		log.Printf("error sending email: %v", err)
		return nil, status.Error(codes.Internal, "error sending email")

	}

	// Context errors handling
	if ctx.Err() == context.Canceled {
		log.Printf("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}

	res := &pb.BillingEmailResponse{
		Message: fmt.Sprintf("email successfully sent to %s", user.GetEmail()),
	}
	return res, nil
}

func (server *EmailServer) SendVerificationEmail(ctx context.Context, req *pb.VerificationEmailRequest) (*pb.VerificationEmailResponse, error) {
	email := req.GetEmail()
	name := req.GetName()
	verificationLink := req.GetVerifictionLink()
	log.Printf("received request for email %s", email)

	templateData := domain.VerificationData{
		Name:             name,
		Email:            email,
		VerificationLink: verificationLink,
	}

	err := helpers.ParseAndSend("templates/verification.html", templateData, email, "Account Verification Email")
	if err != nil {
		log.Printf("error sending email: %v", err)
		return nil, status.Error(codes.Internal, "error sending email")
	}

	// Context errors handling
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
