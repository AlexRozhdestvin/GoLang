package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc/proto"
)

// Реализация сервиса MyService
type myServiceServer struct {
	pb.UnimplementedUserServiceServer
}

// Реализация метода GetData
func (s *myServiceServer) GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Пример ответа
	return &pb.GetUserResponse{Id: "1", Name: "Alex", Age: 28}, nil
}

func main() {
	// Создаем новый gRPC сервер
	srv := grpc.NewServer()

	// Регистрируем наш сервис
	pb.RegisterUserServiceServer(srv, &myServiceServer{})

	// Слушаем на порту (например, 50051)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Запускаем сервер
	log.Println("gRPC сервер запущен на порту 50051")
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
