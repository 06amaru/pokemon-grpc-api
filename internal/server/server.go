package server

import (
	"context"
	"database/sql"
	"pokemon-grpc-api/proto"
)

type PokemonServer struct {
	proto.UnimplementedPokemonServiceServer
	db *sql.DB
}

func NewPokemonServer(db *sql.DB) *PokemonServer {
	return &PokemonServer{db: db}
}

func (s *PokemonServer) CreatePokemon(ctx context.Context, req *proto.CreatePokemonRequest) (*proto.Pokemon, error) {
	var id int32
	err := s.db.QueryRowContext(ctx, "INSERT INTO pokemon (name, type, hp) VALUES ($1, $2, $3) RETURNING id",
		req.Name,
		req.Type,
		req.Hp,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &proto.Pokemon{
		Id:   id,
		Name: req.Name,
		Type: req.Type,
		Hp:   req.Hp,
	}, nil
}

func (s *PokemonServer) GetPokemon(ctx context.Context, req *proto.GetPokemonRequest) (*proto.Pokemon, error) {
	var pokemon proto.Pokemon
	err := s.db.QueryRowContext(ctx, "SELECT id, name, type, hp FROM pokemon WHERE id = $1", req.Id).
		Scan(&pokemon.Id, &pokemon.Name, &pokemon.Type, &pokemon.Hp)

	if err != nil {
		return nil, err
	}
	return &pokemon, nil
}

func (s *PokemonServer) UpdatePokemon(ctx context.Context, req *proto.UpdatePokemonRequest) (*proto.Pokemon, error) {
	_, err := s.db.ExecContext(ctx, "UPDATE pokemon SET name = $1, type = $2, hp = $3 WHERE id = $4",
		req.Name,
		req.Type,
		req.Hp,
		req.Id,
	)
	if err != nil {
		return nil, err
	}

	return &proto.Pokemon{Id: req.Id, Name: req.Name, Type: req.Type, Hp: req.Hp}, nil
}

func (s *PokemonServer) DeletePokemon(ctx context.Context, req *proto.DeletePokemonRequest) (*proto.DeletePokemonResponse, error) {
	result, err := s.db.ExecContext(ctx, "DELETE FROM pokemon WHERE id = $1", req.Id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &proto.DeletePokemonResponse{Success: rowsAffected > 0}, nil
}

func (s *PokemonServer) ListPokemon(ctx context.Context, req *proto.ListPokemonRequest) (*proto.ListPokemonResponse, error) {
	offset := (req.PageNumber - 1) * req.PageSize
	rows, err := s.db.QueryContext(ctx, "SELECT id, name, type, hp FROM pokemon LIMIT $1 OFFSET $2",
		req.PageSize,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pokemons []*proto.Pokemon
	for rows.Next() {
		var pokemon proto.Pokemon
		if err = rows.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Type, &pokemon.Hp); err != nil {
			return nil, err
		}
		pokemons = append(pokemons, &pokemon)
	}
	return &proto.ListPokemonResponse{Pokemon: pokemons}, nil
}
