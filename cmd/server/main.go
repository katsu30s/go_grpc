package main

import (
	// (一部抜粋)
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	hellopb "go_grpc/pkg/grpc"
	pokemonpb "go_grpc/pkg/grpc/pokemon"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
	pokemonpb.UnimplementedPokemonServiceServer
}

func NewMyServer() *myServer {
	return &myServer{}
}

func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func (s *myServer) Pokemon(ctx context.Context, req *pokemonpb.PokemonRequest) (*pokemonpb.PokemonResponse, error) {
	return &pokemonpb.PokemonResponse{
			Id:             req.Id,
			Name:           "Pikachu",
			BaseExperience: 112,
			Height:         4,
			IsDefault:      true,
			Order:          1,
			Weight:         60,
	}, nil
}

func main() {
	// 1. 8080番portのListenerを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// 2. gRPCサーバーを作成
	s := grpc.NewServer()

	// 各サービスの実装を登録
	hellopb.RegisterGreetingServiceServer(s, NewMyServer())
	pokemonpb.RegisterPokemonServiceServer(s, NewMyServer())
	reflection.Register(s)

	// 3. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	// 4.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
