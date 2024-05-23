package start

import (
	"net"

	"github.com/dogkerr/mailing-service/m/v2/internal/grpc"
	"github.com/dogkerr/mailing-service/m/v2/service"
	"go.uber.org/zap"
)

func InitHTTPAndGRPC(address string) {
	// gRPC
	listener, err := net.Listen("tcp", address)
	if err != nil {
		zap.L().Fatal("cannot start server: ", zap.Error(err))
	}

	emailServer := service.NewEmailServer()

	go func() {
		err := grpc.RunGRPCServer(emailServer, listener)
		if err != nil {
			zap.L().Fatal("cannot start GRPC Server", zap.Error(err))
		}
	}()
}
