package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/SystemAlgoFund/grpc_package/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedServiceServer
	handlers map[string]func(context.Context, *pb.Request) (*pb.Response, error)
}

func NewServer(handlers map[string]func(context.Context, *pb.Request) (*pb.Response, error)) *Server {
	return &Server{handlers: handlers}
}

func (s *Server) Send(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	if handler, ok := s.handlers[req.Route]; ok {
		return handler(ctx, req)
	}
	return nil, fmt.Errorf("handler not found for method")
}

func (s *Server) Start(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterServiceServer(grpcServer, s)

	log.Printf("gRPC server listening on %s", address)
	return grpcServer.Serve(lis)
}
