package service

import (
	"context"
	"errors"
	"fmt"
	"net"

	rpc "github.com/GermanVor/data-keeper/cmd/storageServer/rpc"
	"github.com/GermanVor/data-keeper/cmd/storageServer/storage"
	pbDatakeeper "github.com/GermanVor/data-keeper/proto/datakeeper"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	pbUser "github.com/GermanVor/data-keeper/proto/user"
)

type Interface interface {
	Start() error
}

type Impl struct {
	addr        string
	server      *grpc.Server
	serviceImpl *rpc.DatakeeperServiceImpl
}

func (s *Impl) Start() error {
	listen, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	pbDatakeeper.RegisterDataKeeperServer(s.server, s.serviceImpl)

	fmt.Println("Server gRPC started")

	return s.server.Serve(listen)
}

var _ Interface = (*Impl)(nil)

func CheckAccessInterceptor(userClient pbUser.UserClient) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("")
		}

		jwtArr := md.Get("jwt")
		if len(jwtArr) == 0 {
			return nil, errors.New("")
		}

		checkAccessResp, err := userClient.CheckAccess(ctx, &pbUser.CheckAccessRequest{Token: jwtArr[0]})
		if err != nil {
			return nil, errors.New("")
		}

		newCtx := context.WithValue(ctx, "userId", checkAccessResp.UserId)

		return handler(newCtx, req)
	}
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

func Init(addr string, stor storage.Interface, userAddr string) Interface {
	conn, err := grpc.Dial(userAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil
	}

	logrusEntry := logrus.NewEntry(logrusLogger)
	logrusOpts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(customFunc),
	}
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	return &Impl{
		addr: addr,
		server: grpc.NewServer(
			grpc_middleware.WithUnaryServerChain(
				grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
				grpc_logrus.UnaryServerInterceptor(logrusEntry, logrusOpts...),
				CheckAccessInterceptor(pbUser.NewUserClient(conn)),
			),
		),
		serviceImpl: rpc.Init(stor),
	}
}
