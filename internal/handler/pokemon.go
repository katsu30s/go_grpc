package handler

import (
	"context"
	pb "go_grpc/pkg/grpc/pokemon"
)

type pokemonServiceServer struct {
	pb.UnimplementedPokemonServiceServer
}

func (s *pokemonServiceServer) Pokemon(ctx context.Context, req *pb.PokemonRequest) (*pb.PokemonResponse, error) {
	// ここに実際の処理を実装します
	// 例として、固定のレスポンスを返します
	return &pb.PokemonResponse{
			Id:             req.Id,
			Name:           "Pikachu",
			BaseExperience: 112,
			Height:         4,
			IsDefault:      true,
			Order:          1,
			Weight:         60,
	}, nil
}