package rpc

import (
	"context"

	"github.com/charlesonunze/grpc-nats-envoy/user-service/internal/db"
	"github.com/charlesonunze/grpc-nats-envoy/user-service/internal/repo"
	services "github.com/charlesonunze/grpc-nats-envoy/user-service/internal/service"
	"github.com/charlesonunze/grpc-nats-envoy/user-service/pb"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type server struct {
	nc *nats.Conn
}

// New - returns an instance of the UserServiceRPCServer
func New(nc *nats.Conn) pb.UserServiceRPCServer {
	return &server{
		nc: nc,
	}
}

func (s *server) GetService(db *gorm.DB, natsConn *nats.Conn) services.UserService {
	userRepo := repo.New(db)
	return services.New(userRepo, natsConn)
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	db := db.DB
	svc := s.GetService(db, s.nc)

	jwt, err := svc.LoginUser(ctx, req.Body.Name)
	if err != nil {
		return &pb.LoginResponse{}, err
	}

	return &pb.LoginResponse{
		Token: jwt,
	}, nil
}

func (s *server) GetUserBalance(ctx context.Context, req *pb.GetUserBalanceRequest) (*pb.GetUserBalanceResponse, error) {
	db := db.DB
	svc := s.GetService(db, s.nc)

	balance, err := svc.GetUserBalance(ctx, req.Body.Token)

	if err != nil {
		return &pb.GetUserBalanceResponse{}, err
	}

	return &pb.GetUserBalanceResponse{
		Amount: balance,
	}, nil
}
