package server

import (
	"database/sql"
	"grpc_crud_go/internal/handler"
	"grpc_crud_go/internal/service"
	"grpc_crud_go/internal/store"
	pb "grpc_crud_go/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGRPCServer(db *sql.DB) *grpc.Server {
	grpcServer := grpc.NewServer()

	userStore := store.NewUserStore(db)
	userService := service.NewUserService(userStore)
	pb.RegisterUserServiceServer(grpcServer, handler.NewUserHandler(userService))
	reflection.Register(grpcServer)

	return grpcServer
}

func Listen(addr string) (net.Listener, error) {
	return net.Listen("tcp", addr)
}
