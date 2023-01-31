package service

import (
	"net"

	userRPC "github.com/GermanVor/data-keeper/cmd/userStorageServer/rpc"
	"github.com/GermanVor/data-keeper/cmd/userStorageServer/storage"
	pbUser "github.com/GermanVor/data-keeper/proto/user"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Interface interface {
	Start() error
}

type Service struct {
	addr   string
	server *grpc.Server
	impl   *userRPC.UserRPCImpl
}

func (s *Service) Start() error {
	listen, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	pbUser.RegisterUserServer(s.server, s.impl)

	return s.server.Serve(listen)
}

var (
	logrusLogger = logrus.New()
	customFunc   = func(code codes.Code) logrus.Level {
		if code == codes.OK {
			return logrus.InfoLevel
		}
		return logrus.ErrorLevel
	}
)

func Init(addr string, stor storage.Interface) Interface {
	logrusEntry := logrus.NewEntry(logrusLogger)
	logrusOpts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(customFunc),
	}
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	return &Service{
		addr: addr,
		server: grpc.NewServer(
			grpc_middleware.WithUnaryServerChain(
				grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
				grpc_logrus.UnaryServerInterceptor(logrusEntry, logrusOpts...),
			),
		),
		impl: userRPC.Init(stor),
	}
}
